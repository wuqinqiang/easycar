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
	ErrEmpty = errors.New("The Group cannot be empty")
)

// Client client-go
type Client struct {
	//  easycar service uri. the format rules: direct: "ip:port,ip:port" and  discovery:"easycar"
	uri string
	// grpc client conn
	grpcConn *grpc.ClientConn
	// easycarcli client
	easycarCli proto.EasyCarClient
	options    *Options
}

func New(uri string, options ...Option) (client *Client, err error) {
	server := BuildDirectTarget(uri)
	opts := DefaultOptions
	for _, fn := range options {
		fn(opts)
	}
	if opts.isDiscovery {
		server = BuildDiscoveryTarget(uri)
	}

	ctx, cancel := context.WithTimeout(context.Background(), opts.connTimeout)
	defer cancel()

	client = &Client{
		uri:     server,
		options: opts,
	}
	client.easycarCli, err = client.getConn(ctx)
	return
}

func (client *Client) Begin(ctx context.Context) (gid string, err error) {
	resp, err := client.easycarCli.Begin(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.GetGId(), nil
}

func (client *Client) Register(ctx context.Context, gid string, groups []*Group) error {
	if len(groups) == 0 {
		return ErrEmpty
	}
	var req proto.RegisterReq
	req.GId = gid

	for _, group := range groups {
		for _, branch := range group.branches {
			b := branch.Convert()
			b.TranType = consts.ConvertTranTypeToGrpc(group.GetTranType())
			req.Branches = append(req.Branches, b)
		}
	}
	_, err := client.easycarCli.Register(ctx, &req)
	return err
}

func (client *Client) Start(ctx context.Context, gid string) (err error) {
	if client.options.beforeFunc != nil {
		if err = client.options.beforeFunc(ctx); err != nil {
			return
		}
	}
	if client.options.afterFunc != nil {
		defer func() {
			err = client.options.afterFunc(ctx)
		}()
	}
	defer func() {
		if err != nil {
			fmt.Printf("gid:%v Start err:%v\n", gid, err)
			if err = client.Rollback(ctx, gid); err != nil {
				fmt.Printf("gid:%v rollback err:%v\n", gid, err)
				return
			}
			return
		}
		if err = client.Commit(ctx, gid); err != nil {
			fmt.Printf("gid:%v commit err:%v\n", gid, err)
		}
	}()

	var req proto.StartReq
	req.GId = gid

	_, err = client.easycarCli.Start(ctx, &req)

	return err
}

func (client *Client) Commit(ctx context.Context, gid string) error {
	var (
		req proto.CommitReq
	)
	req.GId = gid
	_, err := client.easycarCli.Commit(ctx, &req)
	return err
}

func (client *Client) Rollback(ctx context.Context, gid string) error {
	var (
		req proto.RollBckReq
	)
	req.GId = gid
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
