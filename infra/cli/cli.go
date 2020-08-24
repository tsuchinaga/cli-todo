package cli

import (
	"bufio"
	"fmt"
	"os"

	"gitlab.com/tsuchinaga/cli-todo/infra"
	"gitlab.com/tsuchinaga/cli-todo/infra/store"

	"gitlab.com/tsuchinaga/cli-todo/app/service"
	"gitlab.com/tsuchinaga/cli-todo/app/usecase"

	"gitlab.com/tsuchinaga/cli-todo/adapter"
)

func Run() error {
	newCLI().Run()
	return nil
}

func newCLI() *cli {
	bs := bufio.NewScanner(os.Stdin)
	bs.Split(bufio.ScanLines)
	return &cli{
		bs: bs,
		todoAdapter: adapter.NewTODO(
			usecase.NewTODO(
				service.NewTODO(
					store.NewTODO(),
					infra.NewUUID(),
					infra.NewClock()))),
	}
}

type cli struct {
	bs          *bufio.Scanner
	todoAdapter adapter.TODO
}

func (c *cli) Run() {
	for {
		fmt.Print(">>> ")
		c.bs.Scan()
		cmd := c.bs.Text()
		switch cmd {
		case "help":
			c.help()
		case "list":
			c.list()
		case "create":
			c.create()
		case "delete":
			c.delete()
		case "exit":
			fmt.Println("終了します")
			return
		default:
			fmt.Printf("コマンドがありません(%s)\n", cmd)
		}
	}
}

func (c *cli) help() {
	fmt.Println("コマンド一覧")
	fmt.Println("  help    コマンド一覧")
	fmt.Println("  list    TODO一覧")
	fmt.Println("  create  TODOの追加")
	fmt.Println("  delete  TODOの削除")
	fmt.Println("  exit    終了")
}

func (c *cli) list() {
	todos, err := c.todoAdapter.List()
	if err != nil {
		fmt.Printf("エラーが発生しました(%s)\n", err)
		return
	}

	fmt.Println("TODO一覧(ID, タイトル)")
	for _, todo := range todos {
		fmt.Printf("  %s %s\n", todo.ID, todo.Title)
	}
}

func (c *cli) create() {
	fmt.Println("タイトルを入力してください")
	fmt.Print(">>> ")
	c.bs.Scan()

	title := c.bs.Text()
	id, err := c.todoAdapter.Create(title)
	if err != nil {
		fmt.Printf("追加に失敗しました(タイトル: %s)\n", title)
	}
	fmt.Printf("ID: %sで追加しました\n", id)
}

func (c *cli) delete() {
	fmt.Println("削除するTODOのIDを入力してください")
	fmt.Print(">>> ")
	c.bs.Scan()

	id := c.bs.Text()
	err := c.todoAdapter.Delete(id)
	if err != nil {
		fmt.Printf("削除に失敗しました(ID: %s)\n", id)
	}
	fmt.Printf("ID: %sを削除しました\n", id)
}
