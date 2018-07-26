package word

import (
	"saltlab/foundation/db"
	"saltlab/foundation/xlog"
)

// LoadDesc loads all item descriptions from the DB
func (m *Manager) LoadDesc() error {
	var descs []*struct {
		FilterType int
		Word       string
	}

	_, err := db.BaseSQL().Select("filter_type,word").From("nb_bt_word_filter").OrderBy("sn").Load(&descs)
	if err != nil {
		re := db.NewScanError(err)
		xlog.Errln(re)
		return re
	}

	for _, d := range descs {
		switch d.FilterType {
		case 0:
			m.nameFilter.addFilterWord(d.Word)
			m.chatFilter.addFilterWord(d.Word)
		case 1:
			m.nameFilter.addFilterWord(d.Word)
		case 2:
			m.chatFilter.addFilterWord(d.Word)
		}
	}

	m.nameFilter.addIgnoreChar(" ~!@#$%^&*()_+`-=")
	m.chatFilter.addIgnoreChar(" ~!@#$%^&*()_+`-=")

	return nil
}
