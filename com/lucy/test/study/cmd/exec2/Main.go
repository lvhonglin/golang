package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	root := "C:\\Users\\Administrator\\GolandProjects\\golang\\com\\lucy\\test\\study\\cmd\\exec2"
	command := fmt.Sprintf(`%s`, filepath.Join(root, "guitar_bash.sh"))
	cmd := exec.Command("sh", command)
	//var combinedOutput bytes.Buffer
	//cmd.Stdout = &combinedOutput
	//cmd.Stderr = &combinedOutput
	cmd.Dir = root
	Do(cmd)
	//println(combinedOutput.String())

}

// Do 执行命令，并实时打印输出
func Do(cmd *exec.Cmd) error {
	// 重定向输入
	cmd.Stdin = os.Stdin
	// 打印 stdout 和 stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("get stdout pipe error: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("get stderr pipe error: %v", err)
	}

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("cmd start error: %v", err)
	}
	go scanAndPrint(stdout)
	go scanAndPrint(stderr)
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("cmd wait error: %v", err)
	}

	return nil
}

// scanAndPrint 扫描然后打印输出
func scanAndPrint(r io.ReadCloser) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintln(os.Stderr, string(scanner.Bytes()))
	}
}
