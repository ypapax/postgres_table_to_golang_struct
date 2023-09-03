postgres_table_to_golang_struct converts a file with one postgres `create table` expression to golang struct declaration expression.
# Install
go install
# Usage:
1. crete a file ~/tmp/table1.sql with postgres create table expression, for example:
```
CREATE TABLE IF NOT EXISTS income_statement
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
);
```

2. run: 
`postgres_table_to_golang_struct ~/tmp/table1.sql`

3. then you can find the following golang struct generated code in file ~/tmp/table1.sql.golung.generated.struct

```
type IncomeStatement struct {
    tableName struct{} `json:"-" pg:"income_statement,discard_unknown_columns"`
    ID int `json:"id" pg:"id"`
    Date string `json:"date" pg:"date"`
    Symbol string `json:"symbol" pg:"symbol"`
    ReportedCurrency string `json:"reportedCurrency" pg:"reportedCurrency"`
    Cik string `json:"cik" pg:"cik"`
    FillingDate string `json:"fillingDate" pg:"fillingDate"`
    AcceptedDate string `json:"acceptedDate" pg:"acceptedDate"`
    CalendarYear int `json:"calendarYear" pg:"calendarYear"`
    Period string `json:"period" pg:"period"`
    Revenue float64 `json:"revenue" pg:"revenue"`
    CostOfRevenue float64 `json:"costOfRevenue" pg:"costOfRevenue"`
    GrossProfit float64 `json:"grossProfit" pg:"grossProfit"`
    GrossProfitRatio float64 `json:"grossProfitRatio" pg:"grossProfitRatio"`
    ResearchAndDevelopmentExpenses float64 `json:"researchAndDevelopmentExpenses" pg:"ResearchAndDevelopmentExpenses"`
    GeneralAndAdministrativeExpenses float64 `json:"generalAndAdministrativeExpenses" pg:"GeneralAndAdministrativeExpenses"`
    SellingAndMarketingExpenses float64 `json:"sellingAndMarketingExpenses" pg:"SellingAndMarketingExpenses"`
    SellingGeneralAndAdministrativeExpenses float64 `json:"sellingGeneralAndAdministrativeExpenses" pg:"SellingGeneralAndAdministrativeExpenses"`
    OtherExpenses float64 `json:"otherExpenses" pg:"otherExpenses"`
    OperatingExpenses float64 `json:"operatingExpenses" pg:"operatingExpenses"`
    CostAndExpenses float64 `json:"costAndExpenses" pg:"costAndExpenses"`
    InterestExpense float64 `json:"interestExpense" pg:"interestExpense"`
    DepreciationAndAmortization float64 `json:"depreciationAndAmortization" pg:"depreciationAndAmortization"`
    Ebitda float64 `json:"ebitda" pg:"EBITDA"`
    Ebitdaratio float64 `json:"ebitdaratio" pg:"EBITDARatio"`
    OperatingIncome float64 `json:"operatingIncome" pg:"operatingIncome"`
    OperatingIncomeRatio float64 `json:"operatingIncomeRatio" pg:"operatingIncomeRatio"`
    TotalOtherIncomeExpensesNet float64 `json:"totalOtherIncomeExpensesNet" pg:"totalOtherIncomeExpensesNet"`
    IncomeBeforeTax float64 `json:"incomeBeforeTax" pg:"incomeBeforeTax"`
    IncomeBeforeTaxRatio float64 `json:"incomeBeforeTaxRatio" pg:"incomeBeforeTaxRatio"`
    IncomeTaxExpense float64 `json:"incomeTaxExpense" pg:"incomeTaxExpense"`
    NetIncome float64 `json:"netIncome" pg:"netIncome"`
    NetIncomeRatio float64 `json:"netIncomeRatio" pg:"netIncomeRatio"`
    Eps float64 `json:"eps" pg:"EPS"`
    Epsdiluted float64 `json:"epsdiluted" pg:"EPSDiluted"`
    WeightedAverageShsOut float64 `json:"weightedAverageShsOut" pg:"weightedAverageShsOut"`
    WeightedAverageShsOutDil float64 `json:"weightedAverageShsOutDil" pg:"weightedAverageShsOutDil"`
    Link string `json:"link" pg:"link"`
    FinalLink string `json:"finalLink" pg:"finalLink"`
    InterestIncome float64 `json:"interestIncome" pg:"interestIncome"`
    CreatedAt time.Time `json:"createdAt" pg:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" pg:"updated_at"`
    DateTimestamp time.Time `json:"dateTimestamp" pg:"date_timestamp"`
    Ticker string `json:"ticker" pg:"ticker"`
    FillingDateTimestamp time.Time `json:"fillingDateTimestamp" pg:"fillingDate_timestamp"`
    RevenueChange float64 `json:"revenueChange" pg:"revenueChange"`
    GrossProfitChange float64 `json:"grossProfitChange" pg:"grossProfitChange"`
    OperatingIncomeChange float64 `json:"operatingIncomeChange" pg:"operatingIncomeChange"`
    NetIncomeChange float64 `json:"netIncomeChange" pg:"netIncomeChange"`
    EbitdaChange float64 `json:"ebitdaChange" pg:"ebitdaChange"`
    RawData string `json:"rawData" pg:"raw_data"`
    URL string `json:"url" pg:"url"`
    IncomeStatementApiId int `json:"incomeStatementApiId" pg:"income_statement_api_id"`
}
```