package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	config "github.com/Mr-Herod/CloudGamingDemo/Gaming/config"
	record "github.com/Mr-Herod/CloudGamingDemo/Record/record"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var dirPath string
var WinPos string
var ffmpegPath string

func InitGame() {
	cfg := config.ServiceConf
	dirPath = cfg.GameImgPath
	WinPos = cfg.GameWinPos
	ffmpegPath = cfg.FFmpegPath
}

func StartGame(ctx context.Context, port string) {
	cmd := exec.Command("/bin/bash", "-c",
		"ln -sf "+dirPath+"1.jpg"+" "+dirPath+port+".jpg")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := execCmd(cmd, false)
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	fmt.Printf("ln -s ok\n")

	cmd = exec.Command("/bin/bash", "-c",
		ffmpegPath+" -loop 1 -f image2 -i "+
			dirPath+"'"+port+".jpg' -r 2 -vsync 1 -async 1 "+
			"-f lavfi -vcodec libvpx -cpu-used 5 -deadline 1 "+
			"-g 10 -error-resilient 1 -auto-alt-ref 1 -speed 1 "+
			"-f rtp 'rtp://localhost:"+port+"?pkt_size=1200' > "+
			dirPath+"ffmpeg.log &")
	err = execCmd(cmd, false)
	cmdPid := cmd.Process.Pid + 1 //查看命令pid
	fmt.Println("cmdPid:", cmdPid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ffmpeg ok\n")
}

func execCmd(cmd *exec.Cmd, debug bool) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		if debug {
			log.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
		}
	}
	return err
}

func KillGame(ctx context.Context, port string) {
	saveRecord(ctx)
	cmd := exec.Command("/bin/bash", "-c",
		"kill $(ps aux | grep '"+ffmpegPath+"' | "+
			"grep 'squares/"+port+".jpg' | tr -s ' '| cut -d ' ' -f 2)")
	err := execCmd(cmd, false)
	if err != nil {
		fmt.Printf("Kill game error:%v", err)
	}
	fmt.Printf("Kill game success:%v", err)
}

func saveRecord(ctx context.Context) {
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Record")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	cli := record.NewRecordServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rsp, err := cli.SaveRecord(ctx, &record.SaveRecordReq{
		Username: "herod",
		Gamename: "推箱子",
		Score:    0,
	})
	if err != nil || rsp.ErrCode != 0 {
		log.Printf("UserLogin error: %v", err)
		return
	}
	log.Printf("rsp: %v", rsp)
}

func Move(nowPos int, move string, port string) (int, bool) {
	fmt.Printf("receive command:%s\n", move)
	nextPosi := NextPos(nowPos, move)
	nextPos := strconv.Itoa(nextPosi)
	isWin := (nextPos == WinPos)
	cmd := exec.Command("/bin/bash", "-c", "ln -sf "+dirPath+nextPos+".jpg"+" "+dirPath+port+".jpg")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return nowPos, isWin
	}
	fmt.Printf("NowPos:%v \n", nextPosi)
	return nextPosi, isWin
}

func NextPos(nowPos int, move string) int {
	if nowPos == 8 {
		return 1
	}
	if nowPos == 1 && move == "a" {
		return 2
	}
	if nowPos == 1 && move == "w" {
		return 4
	}
	if nowPos == 2 && move == "w" {
		return 3
	}
	if nowPos == 2 && move == "d" {
		return 1
	}
	if nowPos == 3 && move == "d" {
		return 4
	}
	if nowPos == 3 && move == "w" {
		return 5
	}
	if nowPos == 3 && move == "s" {
		return 2
	}
	if nowPos == 4 && move == "a" {
		return 3
	}
	if nowPos == 4 && move == "w" {
		return 6
	}
	if nowPos == 4 && move == "s" {
		return 1
	}
	if nowPos == 5 && move == "s" {
		return 3
	}
	if nowPos == 5 && move == "d" {
		return 6
	}
	if nowPos == 6 && move == "a" {
		return 5
	}
	if nowPos == 6 && move == "d" {
		return 7
	}
	if nowPos == 6 && move == "s" {
		return 4
	}
	if nowPos == 7 && move == "d" {
		return 8
	}
	if nowPos == 7 && move == "a" {
		return 13
	}
	if nowPos == 9 && move == "a" {
		return 10
	}
	if nowPos == 9 && move == "w" {
		return 14
	}
	if nowPos == 10 && move == "w" {
		return 11
	}
	if nowPos == 10 && move == "d" {
		return 9
	}
	if nowPos == 11 && move == "d" {
		return 14
	}
	if nowPos == 11 && move == "w" {
		return 12
	}
	if nowPos == 11 && move == "s" {
		return 10
	}
	if nowPos == 12 && move == "d" {
		return 13
	}
	if nowPos == 12 && move == "s" {
		return 11
	}
	if nowPos == 13 && move == "s" {
		return 14
	}
	if nowPos == 13 && move == "d" {
		return 7
	}
	if nowPos == 13 && move == "a" {
		return 12
	}
	if nowPos == 14 && move == "a" {
		return 11
	}
	if nowPos == 14 && move == "w" {
		return 13
	}
	if nowPos == 14 && move == "s" {
		return 9
	}
	return nowPos
}
