package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPostgresTableToGolangStruct(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inp         string
		expFileName string
	}
	cases := []testCase{{inp: `CREATE TABLE IF NOT EXISTS income_statement
(
    id         bigserial NOT NULL,
    -- from :
    date text not null,
    symbol text not null,
    reportedCurrency text not null,
    cik text not null,
    fillingDate text not null,
    acceptedDate text,
    calendarYear int,
    period text not null,
    revenue numeric not null,
    costOfRevenue numeric not null,
    grossProfit numeric not null,
    grossProfitRatio numeric not null,
    ResearchAndDevelopmentExpenses numeric,
    GeneralAndAdministrativeExpenses numeric not null,
    SellingAndMarketingExpenses numeric not null,
    SellingGeneralAndAdministrativeExpenses numeric not null,
    otherExpenses numeric not null,
    operatingExpenses numeric not null,
    costAndExpenses numeric not null,
    interestExpense numeric not null,
    depreciationAndAmortization numeric not null,
    EBITDA numeric not null,
    EBITDARatio numeric not null,
    operatingIncome numeric not null,
    operatingIncomeRatio numeric not null,
    totalOtherIncomeExpensesNet numeric not null,
    incomeBeforeTax numeric not null,
    incomeBeforeTaxRatio numeric not null,
    incomeTaxExpense numeric not null,
    netIncome numeric not null,
    netIncomeRatio numeric not null,
    EPS numeric not null,
    EPSDiluted numeric not null,
    weightedAverageShsOut numeric not null,
    weightedAverageShsOutDil numeric not null,
    link text,
    finalLink text,
    interestIncome numeric not null,
-- calculated:
    created_at timestamp default now(),
    updated_at timestamp default now(),
    date_timestamp timestamp not null,
    ticker text,
    fillingDate_timestamp timestamp not null,

    revenueChange numeric,
    grossProfitChange numeric,
    operatingIncomeChange numeric,
    netIncomeChange numeric,
    ebitdaChange numeric,
    raw_data text,
    url text,

    income_statement_api_id int,
    CONSTRAINT income_statement_pk PRIMARY KEY (id)
);`, expFileName: `0.expected`}}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			t.Parallel()
			r := require.New(t)
			act, err := postgresTableToGolangStruct(c.inp)
			r.NoError(err)
			b, err := os.ReadFile(c.expFileName)
			r.NoError(err)
			r.Equal(string(b), act)
		})
	}
}

func TestPostgresColumnNameToGolangFieldName(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inp string
		exp string
	}
	cases := []testCase{{inp: `income_statement_api_id`, exp: `IncomeStatementApiId`}}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			t.Parallel()
			r := require.New(t)
			act := postgresColumnNameToGolangFieldName(c.inp)
			r.Equal(c.exp, act)
		})
	}
}
