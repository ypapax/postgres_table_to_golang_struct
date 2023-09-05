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
    ReportedCurrency string `json:"reportedCurrency" pg:"reportedcurrency"`
    Cik string `json:"cik" pg:"cik"`
    FillingDate string `json:"fillingDate" pg:"fillingdate"`
    AcceptedDate string `json:"acceptedDate" pg:"accepteddate"`
    CalendarYear int `json:"calendarYear" pg:"calendaryear"`
    Period string `json:"period" pg:"period"`
    Revenue float64 `json:"revenue" pg:"revenue"`
    CostOfRevenue float64 `json:"costOfRevenue" pg:"costofrevenue"`
    GrossProfit float64 `json:"grossProfit" pg:"grossprofit"`
    GrossProfitRatio float64 `json:"grossProfitRatio" pg:"grossprofitratio"`
    ResearchAndDevelopmentExpenses float64 `json:"researchAndDevelopmentExpenses" pg:"researchanddevelopmentexpenses"`
    GeneralAndAdministrativeExpenses float64 `json:"generalAndAdministrativeExpenses" pg:"generalandadministrativeexpenses"`
    SellingAndMarketingExpenses float64 `json:"sellingAndMarketingExpenses" pg:"sellingandmarketingexpenses"`
    SellingGeneralAndAdministrativeExpenses float64 `json:"sellingGeneralAndAdministrativeExpenses" pg:"sellinggeneralandadministrativeexpenses"`
    OtherExpenses float64 `json:"otherExpenses" pg:"otherexpenses"`
    OperatingExpenses float64 `json:"operatingExpenses" pg:"operatingexpenses"`
    CostAndExpenses float64 `json:"costAndExpenses" pg:"costandexpenses"`
    InterestExpense float64 `json:"interestExpense" pg:"interestexpense"`
    DepreciationAndAmortization float64 `json:"depreciationAndAmortization" pg:"depreciationandamortization"`
    Ebitda float64 `json:"ebitda" pg:"ebitda"`
    Ebitdaratio float64 `json:"ebitdaratio" pg:"ebitdaratio"`
    OperatingIncome float64 `json:"operatingIncome" pg:"operatingincome"`
    OperatingIncomeRatio float64 `json:"operatingIncomeRatio" pg:"operatingincomeratio"`
    TotalOtherIncomeExpensesNet float64 `json:"totalOtherIncomeExpensesNet" pg:"totalotherincomeexpensesnet"`
    IncomeBeforeTax float64 `json:"incomeBeforeTax" pg:"incomebeforetax"`
    IncomeBeforeTaxRatio float64 `json:"incomeBeforeTaxRatio" pg:"incomebeforetaxratio"`
    IncomeTaxExpense float64 `json:"incomeTaxExpense" pg:"incometaxexpense"`
    NetIncome float64 `json:"netIncome" pg:"netincome"`
    NetIncomeRatio float64 `json:"netIncomeRatio" pg:"netincomeratio"`
    Eps float64 `json:"eps" pg:"eps"`
    Epsdiluted float64 `json:"epsdiluted" pg:"epsdiluted"`
    WeightedAverageShsOut float64 `json:"weightedAverageShsOut" pg:"weightedaverageshsout"`
    WeightedAverageShsOutDil float64 `json:"weightedAverageShsOutDil" pg:"weightedaverageshsoutdil"`
    Link string `json:"link" pg:"link"`
    FinalLink string `json:"finalLink" pg:"finallink"`
    InterestIncome float64 `json:"interestIncome" pg:"interestincome"`
    CreatedAt time.Time `json:"createdAt" pg:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" pg:"updated_at"`
    DateTimestamp time.Time `json:"dateTimestamp" pg:"date_timestamp"`
    Ticker string `json:"ticker" pg:"ticker"`
    FillingDateTimestamp time.Time `json:"fillingDateTimestamp" pg:"fillingdate_timestamp"`
    RevenueChange float64 `json:"revenueChange" pg:"revenuechange"`
    GrossProfitChange float64 `json:"grossProfitChange" pg:"grossprofitchange"`
    OperatingIncomeChange float64 `json:"operatingIncomeChange" pg:"operatingincomechange"`
    NetIncomeChange float64 `json:"netIncomeChange" pg:"netincomechange"`
    EbitdaChange float64 `json:"ebitdaChange" pg:"ebitdachange"`
    RawData string `json:"rawData" pg:"raw_data"`
    URL string `json:"url" pg:"url"`
    IncomeStatementApiId int `json:"incomeStatementApiId" pg:"income_statement_api_id"`
}
```