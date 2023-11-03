/*
 * Copyright (C) 2023 Ahton
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package notify

import (
	"context"
	"github.com/Ahton89/vacancies_scrapper/internal/configuration"
	"github.com/Ahton89/vacancies_scrapper/internal/worker/types"
	"math/rand"
	"time"
)

var headers = []string{
	":hidog: Новая вакансия детектед",
	":partyparrot: А у нас тут новая вакансия",
	":tada: Ой, а что тут у нас? ВАКАНСИЯ!",
	":eyes: Воу, воу, смотри что тут",
	":wave: Привет, я тут новую вакансию нашел",
	":not_bad: Смотри какая вакансия",
	":cool-doge: Вакансия, вакансия, вакансия",
	":pepe_ok: Вакансия для твоего друга",
	":kolya-parrot: Коля одобряет эту вакансию",
	":robot_face: Сам бы забрал, но я бот...",
	":roman_pleasure: 0 дней без вакансий",
}

func (n *notifier) RandomHeader() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return headers[rand.Intn(len(headers))]
}

var colors = []string{
	"#B2F347",
	"#75FB4C",
	"#FA9B73",
	"#73E1FA",
	"#7388FA",
	"#AE73FA",
	"#FA73AB",
	"#BEFA73",
}

func (n *notifier) RandomColor() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return colors[rand.Intn(len(colors))]
}

type notifier struct {
	config configuration.Configuration
}

type Notifier interface {
	Notify(ctx context.Context, vacancies []types.VacancyInfo) error
	WelcomeMessage(ctx context.Context, vacanciesCount int) error
}
