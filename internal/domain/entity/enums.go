package entity

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type UserPlan string

const (
	FreePlan    UserPlan = "free"
	PremiumPlan UserPlan = "premium"
	DiamondPlan UserPlan = "diamond"
)

type BankAccount string

const (
	NuBank      BankAccount = "nubank"
	C6Bank      BankAccount = "c6_bank"
	MercadoPago BankAccount = "mercado_pago"
	PagBank     BankAccount = "pag_bank"
	PicPay      BankAccount = "picpay"
	InterBank   BankAccount = "banco_inter"
	OthersBank  BankAccount = "outros"
)

type AccountType string

const (
	Checking      AccountType = "checking"
	Saving        AccountType = "savings"
	CreditCard    AccountType = "credit_card"
	DigitalWallet AccountType = "digital_wallet"
)

type CategoryType string

const (
	Income  CategoryType = "income"
	Expense CategoryType = "expense"
)

type PaymentMethod string

const (
	CardCredit    PaymentMethod = "credit_card"
	CardDebit     PaymentMethod = "debit_card"
	Pix           PaymentMethod = "pix"
	Cash          PaymentMethod = "cash"
	OthersPayment PaymentMethod = "others"
)

type InvestmentType string

const (
	Caixinha      InvestmentType = "caixinha"
	Crypto        InvestmentType = "crypto"
	Poupanca      InvestmentType = "poupanca"
	RendaFixa     InvestmentType = "renda_fixa"
	RendaVariavel InvestmentType = "renda_variavel"
	Outros        InvestmentType = "outros"
)