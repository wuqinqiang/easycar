package client

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/pkg/errors"
	_ "github.com/wuqinqiang/easycar/core/resolver/direct"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	EmptyGroup = errors.New("The Group cannot be empty")
)

// Client easycar client
type Client struct {
	//  easycar service uri. the format rules: direct: "ip:port,ip:port" and  discovery:"easycar"
	uri string
	// currentLevel for set branches currentLevel
	currentLevel consts.Level
	// groups a group of branches
	groups []*Group
	//gid global of one distributed transaction
	gid string
	// grpc client conn
	grpcConn *grpc.ClientConn
	// easycarcli client
	easycarCli proto.EasyCarClient
	options    *Options
}

func New(uri string, options ...Option) (client *Client, err error) {
	opts := DefaultOptions

	for _, fn := range options {
		fn(opts)
	}
	ctx, cancel := context.WithTimeout(context.Background(), opts.connTimeout)
	defer cancel()

	client = &Client{
		uri:          uri,
		currentLevel: 1,
		options:      opts,
	}
	client.easycarCli, err = client.getConn(ctx)
	return
}

func (client *Client) AddGroup(skip bool, groups ...*Group) *Client {
	if skip {
		client.currentLevel++
	}
	for _, group := range groups {
		group.SetLevel(client.currentLevel)
		client.groups = append(client.groups, group)
	}
	return client
}

func (client *Client) SetGid(gid string) *Client {
	client.setGid(gid)
	return client
}

func (client *Client) setGid(gid string) {
	client.gid = gid
}

func (client *Client) Begin(ctx context.Context) (gid string, err error) {
	if err != nil {
		return "", err
	}
	resp, err := client.easycarCli.Begin(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	client.setGid(resp.GetGId())
	return client.gid, nil
}

func (client *Client) Register(ctx context.Context) error {
	if len(client.groups) == 0 {
		return EmptyGroup
	}
	var (
		req proto.RegisterReq
	)
	req.GId = client.gid
	for _, group := range client.groups {
		for _, branch := range group.branches {
			b := branch.Convert()
			b.TranType = consts.ConvertTranTypeToGrpc(group.GetTranType())
			req.Branches = append(req.Branches, b)
		}
	}
	_, err := client.easycarCli.Register(ctx, &req)
	return err
}

func (client *Client) Start(ctx context.Context) (err error) {
	var (
		req proto.StartReq
	)
	req.GId = client.gid
	defer func() {
		if err != nil {
			fmt.Printf("gid:%v Start err:%v\n", client.gid, err)
			if err = client.rollback(ctx); err != nil {
				fmt.Printf("gid:%v rollback err:%v\n", client.gid, err)
				return
			}
			return
		}
		if err = client.commit(ctx); err != nil {
			fmt.Printf("gid:%v commit err:%v\n", client.gid, err)
		}
	}()

	_, err = client.easycarCli.Start(ctx, &req)
	return err
}

func (client *Client) commit(ctx context.Context) error {
	var (
		req proto.CommitReq
	)
	req.GId = client.gid
	_, err := client.easycarCli.Commit(ctx, &req)
	return err
}

func (client *Client) rollback(ctx context.Context) error {
	var (
		req proto.RollBckReq
	)
	req.GId = client.gid
	_, err := client.easycarCli.Rollback(ctx, &req)
	return err
}

func (client *Client) getConn(ctx context.Context) (cli proto.EasyCarClient, err error) {
	if client.easycarCli != nil {
		return client.easycarCli, nil
	}
	options := client.options.dailOpts
	creds := insecure.NewCredentials()
	if client.options.tls != nil {
		creds = credentials.NewTLS(client.options.tls)
	}

	options = append(options, grpc.WithTransportCredentials(creds))
	options = append(options, grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"easycarBalancer":{}}]}`))

	conn, err := grpc.DialContext(ctx, client.uri,
		options...)
	if err != nil {
		return nil, err
	}

	client.grpcConn = conn
	cli = proto.NewEasyCarClient(client.grpcConn)
	client.easycarCli = cli
	return
}

func (client *Client) Close(ctx context.Context) error {
	return client.grpcConn.Close()
}
