package main

func (m *mapgame)Init(level []LevelInfo){
	for _, level := range level {
		m.level = append(m.level, level)
	}
}