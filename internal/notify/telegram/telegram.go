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

package telegram

import (
	"context"
	"github.com/Ahton89/vacancies_scrapper/internal/configuration"
	"github.com/Ahton89/vacancies_scrapper/internal/worker/types"
)

func New(config *configuration.Configuration) Telegram {
	return &telegram{
		config: config,
	}
}

func (t *telegram) Notify(ctx context.Context, vacancies []types.VacancyInfo) (err error) {
	return nil
}

func (t *telegram) WelcomeMessage(ctx context.Context, vacanciesCount int) error {
	return nil
}
