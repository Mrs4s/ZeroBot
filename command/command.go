package command

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

type Command struct {
	Name    string
	Aliases []string

	Short string
	Long  string

	Run func(ctx *Ctx)

	commands []*Command
}

// AddRootCommand 添加根指令
func AddRootCommand(cmd *Command) *zero.Matcher {
	m := zero.Matcher{
		Temp:     false,
		Block:    false,
		Priority: 0,
		Type:     zero.Type("message"),
		Rules:    []zero.Rule{zero.CommandRule(append(cmd.Aliases, cmd.Name)...)},
		Handler: func(ctx *zero.Ctx) {
			c := &Ctx{
				Ctx:  *ctx,
				Cmd:  cmd,
				Args: Parse(ctx.State["args"].(string)),
			}
			cmd.execute(c)
		},
		Engine: nil,
	}
	return zero.StoreMatcher(&m)
}

// AddCommand adds one or more commands to this parent command.
func (c *Command) AddCommand(cmds ...*Command) {
	// TODO(wdvxdr): check the child command name conflict.
	for i, x := range cmds {
		if cmds[i] == c {
			panic("Command can't be a child of itself")
		}
		// cmds[i].parent = c
		c.commands = append(c.commands, x)
	}
}

func (c *Command) execute(ctx *Ctx) {
	// deal the child command
	if len(ctx.Args) >= 1 && !isFlag(ctx.Args[0]) {
		first := ctx.Args[0]
		for _, child := range ctx.Cmd.commands {
			if child.match(first) {
				newCtx := &Ctx{
					Ctx:  ctx.Ctx,
					Cmd:  child,
					Args: ctx.Args[1:],
				}
				child.execute(newCtx)
				return
			}
		}
	}

	// TODO(wdvxdr): deal with the flags
	c.Run(ctx)
}

func (c *Command) match(name string) bool {
	ok := name == c.Name
	for _, alias := range c.Aliases {
		ok = ok || alias == name
	}
	return ok
}

func isFlag(s string) bool { return s[0] == '-' }