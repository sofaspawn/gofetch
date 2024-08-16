package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zcalusic/sysinfo"
)

type model struct{
    info ShowInfo
    cursor int
}

type ShowInfo struct{
    Username string
    Hostname string
    Os_architecture string
    Kernel_architecture string
    Kernel_name string
    Cpu_model string
    Cpu_cores uint
    Cpu_threads uint
    Network_macaddress string
}

func initialModel(i ShowInfo) model {
    return model{
        info: i,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < 9 {
                m.cursor++
            }

        }
    }

    return m, nil
} 

func (m model) View() string {

    tea.ClearScreen()

    s := fmt.Sprintf("%s@%s\n\n", m.info.Username, m.info.Hostname)

    s += "\nPress q to quit.\n"

    return s
}

func (i ShowInfo)render(){
    p := tea.NewProgram(initialModel(i))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

func main(){
    current, err := user.Current();

    if err!=nil{
        log.Fatal(err)
    }

    if current.Uid!="1000"{
        log.Fatal("You're not allowed")
    }

    var si sysinfo.SysInfo

    si.GetSysInfo()

    info := ShowInfo{
        Username:current.Username,
        Hostname:si.Node.Hostname,
        Os_architecture: si.OS.Architecture,
        Kernel_name: si.Kernel.Release,
        Kernel_architecture: si.Kernel.Architecture,
        Cpu_model: si.CPU.Model,
        Cpu_cores: si.CPU.Cores,
        Cpu_threads: si.CPU.Threads,
        Network_macaddress: si.Network[0].MACAddress,
    }

    info.render()

    /*
    - username@hostname
    - os.name
    - os.architecture
    - kernel.release
    - cpu.model
    - cpu.cores + cpu.threads
    - network[0].macadress
    */
}
