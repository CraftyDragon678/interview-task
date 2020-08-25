package framework

type (
	// Command executes
	Command func(*Context)

	// CommandStruct has data of command
	CommandStruct struct {
		name    string
		command Command
		aliases []string
	}

	// CommandHandler handler of command
	CommandHandler struct {
		cmds []CommandStruct
	}
)

// CmdHandler handler of command
var CmdHandler *CommandHandler

// NewCommandHandler makes new command handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make([]CommandStruct, 0)}
}

// GetCmds returns cmds of handler
func (handler CommandHandler) GetCmds() []CommandStruct {
	return handler.cmds
}

// Get returns cmd by name
func (handler CommandHandler) Get(name string) (*Command, bool) {
	for _, cmd := range handler.cmds {
		if cmd.name == name {
			return &cmd.command, true
		}
		for _, alias := range cmd.aliases {
			if alias == name {
				return &cmd.command, true
			}
		}
	}
	return nil, false
}

// Register command
func (handler *CommandHandler) Register(name string, command Command, aliases []string) {
	cmdstruct := CommandStruct{name, command, aliases}
	handler.cmds = append(handler.cmds, cmdstruct)
}

// GetName returns name of command
func (command CommandStruct) GetName() string {
	return command.name
}
