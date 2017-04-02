package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

const EmAndamento = "Em Andamento"
const Encerrada = "Encerrada"

type Team struct {
	Penalti interface{} `json:"penalti"`
	Placar  int         `json:"placar"`
	Sigla   string      `json:"sigla"`
	Escudo  string      `json:"escudo"`
	Nome    string      `json:"nome"`
}

type Game struct {
	TransmissaoID   int         `json:"transmissao_id"`
	Status          string      `json:"status"`
	FaseRodada      string      `json:"fase_rodada"`
	DiaSemana       string      `json:"dia_semana"`
	TimeCasa        Team        `json:"time_casa"`
	VideoAoVivo     string      `json:"video_ao_vivo"`
	TipoTransmissao string      `json:"tipo_transmissao"`
	Hora            string      `json:"hora"`
	NomeCampeonato  string      `json:"nome_campeonato"`
	URL             interface{} `json:"url"`
	TimeVisitante   Team        `json:"time_visitante"`
	Localizacao     string      `json:"localizacao"`
	Data            string      `json:"data"`
	ID              int         `json:"id"`
	Date            time.Time
}

func (ga *Game) Print() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	hiGreen := color.New(color.FgHiGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	if ga.Status == EmAndamento {
		fmt.Print(hiGreen(ga.Status) + " ")
	} else {
		fmt.Print(green(ga.Status) + "(" + ga.Data + " " + ga.Hora + ") ")
	}
	fmt.Print(blue(ga.TimeCasa.Nome) + " ")
	if ga.Status == EmAndamento {
		fmt.Print(red(strconv.Itoa(ga.TimeCasa.Placar)+" X "+strconv.Itoa(ga.TimeVisitante.Placar)) + " ")
	} else {
		fmt.Print(red("X "))
	}
	fmt.Print(yellow(ga.TimeVisitante.Nome) + " ")
	fmt.Println(magenta(ga.NomeCampeonato))
}

type CentralResponse struct {
	DataHoje string `json:"data_hoje"`
	Games    []Game `json:"jogos"`
}

type BFData interface {
	End() ([]Game, error)
	Team(string) BFData
	Championship(string) BFData
	Status(string) BFData
	Date(string) BFData
	DayOfWeek(string) BFData
	Phase(string) BFData
	Location(string) BFData
	Hour(string) BFData
	Today(bool) BFData
	Upcoming(bool) BFData
}

func NewBFData() BFData {
	return &GloboEsporte{
		URL: "http://globoesporte.globo.com/temporeal/futebol/central.json",
	}
}

type GloboEsporte struct {
	URL          string
	team         string
	championship string
	status       string
	date         string
	dayOfWeek    string
	phase        string
	location     string
	hour         string
	today        bool
	upcoming     bool
}

func (g *GloboEsporte) get() (CentralResponse, error) {
	var resp CentralResponse

	_, body, errs := gorequest.New().Get(g.URL).End()
	if errs != nil {
		return resp, errs[0]
	}

	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type ByDate []Game

func (d ByDate) Len() int {
	return len(d)
}

func (d ByDate) Less(i, j int) bool {
	return d[i].Date.Before(d[j].Date)
}

func (d ByDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (g *GloboEsporte) End() ([]Game, error) {
	resp, err := g.get()
	if err != nil {
		return resp.Games, err
	}

	var games []Game
	for _, elem := range resp.Games {
		if g.team != "" && elem.TimeCasa.Nome != g.team && elem.TimeVisitante.Nome != g.team {
			continue
		}

		if g.championship != "" && elem.NomeCampeonato != g.championship {
			continue
		}
		if g.status != "" && elem.Status != g.status {
			continue
		}
		if g.date != "" && elem.Data != g.date {
			continue
		}
		if g.dayOfWeek != "" && elem.DiaSemana != g.dayOfWeek {
			continue
		}
		if g.phase != "" && elem.FaseRodada != g.phase {
			continue
		}
		if g.location != "" && elem.Localizacao != g.location {
			continue
		}
		if g.hour != "" && elem.Hora != g.hour {
			continue
		}
		if g.today == true && elem.Data != resp.DataHoje {
			continue
		}
		if g.upcoming == true && elem.Status == Encerrada {
			continue
		}
		elem.Date, err = time.Parse("02/01/2006 15:04", elem.Data+" "+elem.Hora)
		if err != nil {
			return resp.Games, err
		}
		games = append(games, elem)
	}

	sort.Sort(ByDate(games))

	return games, nil
}

func (g *GloboEsporte) Team(team string) BFData {
	g.team = team
	return g
}

func (g *GloboEsporte) Championship(championship string) BFData {
	g.championship = championship
	return g
}

func (g *GloboEsporte) Status(status string) BFData {
	g.status = status
	return g
}

func (g *GloboEsporte) Date(date string) BFData {
	g.date = date
	return g
}

func (g *GloboEsporte) DayOfWeek(dayOfWeek string) BFData {
	g.dayOfWeek = dayOfWeek
	return g
}

func (g *GloboEsporte) Phase(phase string) BFData {
	g.phase = phase
	return g
}

func (g *GloboEsporte) Location(location string) BFData {
	g.location = location
	return g
}

func (g *GloboEsporte) Hour(hour string) BFData {
	g.hour = hour
	return g
}

func (g *GloboEsporte) Today(today bool) BFData {
	g.today = today
	return g
}

func (g *GloboEsporte) Upcoming(upcoming bool) BFData {
	g.upcoming = upcoming
	return g
}
