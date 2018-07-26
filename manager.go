package word

import (
	"saltlab/foundation/xlog"
	"saltlab/nba/model"
	"sync"
)

// Manager :
type Manager struct {
	nameFilter *WordFilter
	chatFilter *WordFilter
}

var (
	once    sync.Once
	manager *model.Manager
)

// GetManager :
func getManager() *Manager {
	once.Do(initOnce)
	return manager.Get().(*Manager)
}

func initOnce() {
	var err error
	manager, err = model.NewManager(func() model.Loader {
		return &Manager{}
	})
	if err != nil {
		panic(err)
	}
}

// Init ..
func Init() {
	once.Do(initOnce)
}

// Reload :
func Reload() error {
	if err := manager.Reload(); err != nil {
		return err
	}
	return nil
}

// Load :
func (m *Manager) Load() error {
	// clear old data
	m.nameFilter = NewWordFilter()
	m.chatFilter = NewWordFilter()

	// load base data
	if err := m.LoadDesc(); err != nil {
		xlog.Errln("failed Load filter Desc", err)
		return err
	}

	xlog.Logln("word filter descriptions are loaded", xlog.Fields{"nameFilter": m.nameFilter.size, "chatFilter": m.chatFilter.size})
	return nil
}

// GetNameFilter :
func GetNameFilter() *WordFilter {
	m := getManager()

	return m.nameFilter
}

// GetChatFilter :
func GetChatFilter() *WordFilter {
	m := getManager()

	return m.chatFilter
}

// IsInclude check include word filter
func IsInclude(str string) bool {
	m := getManager()
	return m.nameFilter.IsInclude(str)
}

// ToFilter replace word filter
func ToFilter(str string) string {
	m := getManager()
	return m.chatFilter.ToFilter(str)
}
