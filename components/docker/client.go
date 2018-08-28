package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/golang/glog"
	"github.com/docker/docker/api/types"
	"encoding/json"
	"io"
	"os"
	)

const (
	version string = "1.37"
)

func ConnectClient() (interface{}, interface{}) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))

	if err != nil {
		glog.Errorf("docker connect: %v", err)
		panic(err)
	}

	return ctx, cli
}

//docker ps
func ListDocker() string {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		glog.Errorf("docker connect: %v", err)
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic("")
	}
	jsons, err := json.MarshalIndent(containers,""," ")

	return string(jsons)
}

//docker images
func ListImages() string {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		glog.Errorf("docker connect: %v", err)
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	jsons, err := json.Marshal(images)

	return string(jsons)
}

//docker logs
func SelectLog(dockerId string) int64 {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		panic(err)
	}
	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, dockerId, options)
	if err != nil {
		panic(err)
	}

	ints, err := io.Copy(os.Stdout, out)
	return ints
}

// docker pull
func PullDocker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, "alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

//docker start
func StartDocker(strs ...string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))

	defer func() {
		cli.Close()
	}()
	if err != nil {
		panic(err)
	}

	for _, container := range strs {
		conId := container
		if err := cli.ContainerStart(ctx, conId, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}
	}

}

//docker stop
func StopDocker(strs ...string) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		conId := container.ID
		if err := cli.ContainerStop(ctx, conId, nil); err != nil {
			defer func() {
				if err == recover() {
					StartDocker(conId)
				}
			}()
			panic(err)
		}
	}
}

//docker top
func TopDocker(dockerName string) string{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	defer func() {
		cli.Close()
	}()
	if err != nil {
		panic(err)
	}
	pid:=[]string{""}
	dockerPid,err:=cli.ContainerTop(ctx,dockerName,pid)
	if err!=nil{
		glog.Errorf("can not find dockerPid: %v",err)
		return "error"
	}
	jsons, err := json.Marshal(dockerPid)
	return string(jsons)
}
