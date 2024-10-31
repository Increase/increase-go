# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BalanceLookup">BalanceLookup</a>

Methods:

- <code title="post /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNewParams">AccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountUpdateParams">AccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountListParams">AccountListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/balance">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.Balance">Balance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountBalanceParams">AccountBalanceParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BalanceLookup">BalanceLookup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/close">client.Accounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumber">AccountNumber</a>

Methods:

- <code title="post /account_numbers">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberNewParams">AccountNumberNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_numbers/{account_number_id}">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountNumberID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /account_numbers/{account_number_id}">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountNumberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberUpdateParams">AccountNumberUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_numbers">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumberListParams">AccountNumberListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountNumber">AccountNumber</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDetails">CardDetails</a>

Methods:

- <code title="post /cards">client.Cards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_id}">client.Cards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cards/{card_id}">client.Cards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardUpdateParams">CardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards">client.Cards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardListParams">CardListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Card">Card</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_id}/details">client.Cards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardService.Details">Details</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDetails">CardDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardPayments

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>

Methods:

- <code title="get /card_payments/{card_payment_id}">client.CardPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardPaymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_payments">client.CardPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPaymentListParams">CardPaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardPurchaseSupplements

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplement">CardPurchaseSupplement</a>

Methods:

- <code title="get /card_purchase_supplements/{card_purchase_supplement_id}">client.CardPurchaseSupplements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardPurchaseSupplementID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplement">CardPurchaseSupplement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_purchase_supplements">client.CardPurchaseSupplements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplementListParams">CardPurchaseSupplementListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPurchaseSupplement">CardPurchaseSupplement</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardDisputes

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDispute">CardDispute</a>

Methods:

- <code title="post /card_disputes">client.CardDisputes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDisputeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDisputeNewParams">CardDisputeNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_disputes/{card_dispute_id}">client.CardDisputes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDisputeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardDisputeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_disputes">client.CardDisputes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDisputeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDisputeListParams">CardDisputeListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDispute">CardDispute</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhysicalCards

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>

Methods:

- <code title="post /physical_cards">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardNewParams">PhysicalCardNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_cards/{physical_card_id}">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /physical_cards/{physical_card_id}">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardUpdateParams">PhysicalCardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_cards">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardListParams">PhysicalCardListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DigitalCardProfiles

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>

Methods:

- <code title="post /digital_card_profiles">client.DigitalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileNewParams">DigitalCardProfileNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /digital_card_profiles/{digital_card_profile_id}">client.DigitalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, digitalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /digital_card_profiles">client.DigitalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileListParams">DigitalCardProfileListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /digital_card_profiles/{digital_card_profile_id}/archive">client.DigitalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, digitalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /digital_card_profiles/{digital_card_profile_id}/clone">client.DigitalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileService.Clone">Clone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, digitalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfileCloneParams">DigitalCardProfileCloneParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalCardProfile">DigitalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhysicalCardProfiles

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>

Methods:

- <code title="post /physical_card_profiles">client.PhysicalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileNewParams">PhysicalCardProfileNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_card_profiles/{physical_card_profile_id}">client.PhysicalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_card_profiles">client.PhysicalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileListParams">PhysicalCardProfileListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /physical_card_profiles/{physical_card_profile_id}/archive">client.PhysicalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /physical_card_profiles/{physical_card_profile_id}/clone">client.PhysicalCardProfiles.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileService.Clone">Clone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfileCloneParams">PhysicalCardProfileCloneParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCardProfile">PhysicalCardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DigitalWalletTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletToken">DigitalWalletToken</a>

Methods:

- <code title="get /digital_wallet_tokens/{digital_wallet_token_id}">client.DigitalWalletTokens.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletTokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, digitalWalletTokenID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletToken">DigitalWalletToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /digital_wallet_tokens">client.DigitalWalletTokens.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletTokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletTokenListParams">DigitalWalletTokenListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DigitalWalletToken">DigitalWalletToken</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>

Methods:

- <code title="get /transactions/{transaction_id}">client.Transactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#TransactionListParams">TransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PendingTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransaction">PendingTransaction</a>

Methods:

- <code title="get /pending_transactions/{pending_transaction_id}">client.PendingTransactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pendingTransactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransaction">PendingTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pending_transactions">client.PendingTransactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransactionListParams">PendingTransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PendingTransaction">PendingTransaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeclinedTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransaction">DeclinedTransaction</a>

Methods:

- <code title="get /declined_transactions/{declined_transaction_id}">client.DeclinedTransactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, declinedTransactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransaction">DeclinedTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /declined_transactions">client.DeclinedTransactions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransactionListParams">DeclinedTransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DeclinedTransaction">DeclinedTransaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>

Methods:

- <code title="post /account_transfers">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferNewParams">AccountTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_transfers/{account_transfer_id}">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_transfers">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferListParams">AccountTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account_transfers/{account_transfer_id}/approve">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account_transfers/{account_transfer_id}/cancel">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ACHTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>

Methods:

- <code title="post /ach_transfers">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferNewParams">ACHTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_transfers/{ach_transfer_id}">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_transfers">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferListParams">ACHTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ach_transfers/{ach_transfer_id}/approve">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ach_transfers/{ach_transfer_id}/cancel">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ACHPrenotifications

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotification">ACHPrenotification</a>

Methods:

- <code title="post /ach_prenotifications">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotificationNewParams">ACHPrenotificationNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotification">ACHPrenotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_prenotifications/{ach_prenotification_id}">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achPrenotificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotification">ACHPrenotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_prenotifications">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotificationListParams">ACHPrenotificationListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHPrenotification">ACHPrenotification</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundACHTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>

Methods:

- <code title="get /inbound_ach_transfers/{inbound_ach_transfer_id}">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_ach_transfers">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferListParams">InboundACHTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/create_notification_of_change">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferService.NewNotificationOfChange">NewNotificationOfChange</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferNewNotificationOfChangeParams">InboundACHTransferNewNotificationOfChangeParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/decline">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferService.Decline">Decline</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferDeclineParams">InboundACHTransferDeclineParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/transfer_return">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferService.TransferReturn">TransferReturn</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransferTransferReturnParams">InboundACHTransferTransferReturnParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WireTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>

Methods:

- <code title="post /wire_transfers">client.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferNewParams">WireTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_transfers/{wire_transfer_id}">client.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_transfers">client.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferListParams">WireTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /wire_transfers/{wire_transfer_id}/approve">client.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /wire_transfers/{wire_transfer_id}/cancel">client.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundWireTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransfer">InboundWireTransfer</a>

Methods:

- <code title="get /inbound_wire_transfers/{inbound_wire_transfer_id}">client.InboundWireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundWireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransfer">InboundWireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_wire_transfers">client.InboundWireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransferListParams">InboundWireTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransfer">InboundWireTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WireDrawdownRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequest">WireDrawdownRequest</a>

Methods:

- <code title="post /wire_drawdown_requests">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequestNewParams">WireDrawdownRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequest">WireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_drawdown_requests/{wire_drawdown_request_id}">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireDrawdownRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequest">WireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_drawdown_requests">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequestListParams">WireDrawdownRequestListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireDrawdownRequest">WireDrawdownRequest</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundWireDrawdownRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>

Methods:

- <code title="get /inbound_wire_drawdown_requests/{inbound_wire_drawdown_request_id}">client.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundWireDrawdownRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_wire_drawdown_requests">client.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequestListParams">InboundWireDrawdownRequestListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CheckTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>

Methods:

- <code title="post /check_transfers">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferNewParams">CheckTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_transfers/{check_transfer_id}">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_transfers">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferListParams">CheckTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/approve">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/cancel">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/stop_payment">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferService.StopPayment">StopPayment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransferStopPaymentParams">CheckTransferStopPaymentParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundCheckDeposits

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>

Methods:

- <code title="get /inbound_check_deposits/{inbound_check_deposit_id}">client.InboundCheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundCheckDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_check_deposits">client.InboundCheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositListParams">InboundCheckDepositListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_check_deposits/{inbound_check_deposit_id}/decline">client.InboundCheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositService.Decline">Decline</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundCheckDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_check_deposits/{inbound_check_deposit_id}/return">client.InboundCheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositService.Return">Return</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundCheckDepositID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDepositReturnParams">InboundCheckDepositReturnParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RealTimePaymentsTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>

Methods:

- <code title="post /real_time_payments_transfers">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransferNewParams">RealTimePaymentsTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_transfers/{real_time_payments_transfer_id}">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimePaymentsTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_transfers">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransferListParams">RealTimePaymentsTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundRealTimePaymentsTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransfer">InboundRealTimePaymentsTransfer</a>

Methods:

- <code title="get /inbound_real_time_payments_transfers/{inbound_real_time_payments_transfer_id}">client.InboundRealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundRealTimePaymentsTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransfer">InboundRealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_real_time_payments_transfers">client.InboundRealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransferListParams">InboundRealTimePaymentsTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransfer">InboundRealTimePaymentsTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CheckDeposits

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>

Methods:

- <code title="post /check_deposits">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDepositService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDepositNewParams">CheckDepositNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_deposits/{check_deposit_id}">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDepositService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_deposits">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDepositService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDepositListParams">CheckDepositListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lockboxes

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Lockbox">Lockbox</a>

Methods:

- <code title="post /lockboxes">client.Lockboxes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxNewParams">LockboxNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Lockbox">Lockbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /lockboxes/{lockbox_id}">client.Lockboxes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lockboxID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Lockbox">Lockbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /lockboxes/{lockbox_id}">client.Lockboxes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lockboxID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxUpdateParams">LockboxUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Lockbox">Lockbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /lockboxes">client.Lockboxes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#LockboxListParams">LockboxListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Lockbox">Lockbox</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundMailItems

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItem">InboundMailItem</a>

Methods:

- <code title="get /inbound_mail_items/{inbound_mail_item_id}">client.InboundMailItems.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundMailItemID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItem">InboundMailItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_mail_items">client.InboundMailItems.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItemListParams">InboundMailItemListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItem">InboundMailItem</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RoutingNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RoutingNumberListResponse">RoutingNumberListResponse</a>

Methods:

- <code title="get /routing_numbers">client.RoutingNumbers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RoutingNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RoutingNumberListParams">RoutingNumberListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RoutingNumberListResponse">RoutingNumberListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccount">ExternalAccount</a>

Methods:

- <code title="post /external_accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountNewParams">ExternalAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external_accounts/{external_account_id}">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, externalAccountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /external_accounts/{external_account_id}">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, externalAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountUpdateParams">ExternalAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external_accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccountListParams">ExternalAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExternalAccount">ExternalAccount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Entities

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>

Methods:

- <code title="post /entities">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityNewParams">EntityNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityListParams">EntityListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/archive">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/archive_beneficial_owner">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.ArchiveBeneficialOwner">ArchiveBeneficialOwner</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityArchiveBeneficialOwnerParams">EntityArchiveBeneficialOwnerParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/confirm">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityConfirmParams">EntityConfirmParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/create_beneficial_owner">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.NewBeneficialOwner">NewBeneficialOwner</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityNewBeneficialOwnerParams">EntityNewBeneficialOwnerParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/update_address">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.UpdateAddress">UpdateAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityUpdateAddressParams">EntityUpdateAddressParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/update_beneficial_owner_address">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.UpdateBeneficialOwnerAddress">UpdateBeneficialOwnerAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityUpdateBeneficialOwnerAddressParams">EntityUpdateBeneficialOwnerAddressParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/update_industry_code">client.Entities.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityService.UpdateIndustryCode">UpdateIndustryCode</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntityUpdateIndustryCodeParams">EntityUpdateIndustryCodeParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SupplementalDocuments

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntitySupplementalDocument">EntitySupplementalDocument</a>

Methods:

- <code title="post /entity_supplemental_documents">client.SupplementalDocuments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SupplementalDocumentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SupplementalDocumentNewParams">SupplementalDocumentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntitySupplementalDocument">EntitySupplementalDocument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entity_supplemental_documents">client.SupplementalDocuments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SupplementalDocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SupplementalDocumentListParams">SupplementalDocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EntitySupplementalDocument">EntitySupplementalDocument</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Programs

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Program">Program</a>

Methods:

- <code title="get /programs/{program_id}">client.Programs.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProgramService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, programID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Program">Program</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /programs">client.Programs.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProgramService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProgramListParams">ProgramListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Program">Program</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ProofOfAuthorizationRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequest">ProofOfAuthorizationRequest</a>

Methods:

- <code title="get /proof_of_authorization_requests/{proof_of_authorization_request_id}">client.ProofOfAuthorizationRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, proofOfAuthorizationRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequest">ProofOfAuthorizationRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /proof_of_authorization_requests">client.ProofOfAuthorizationRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestListParams">ProofOfAuthorizationRequestListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequest">ProofOfAuthorizationRequest</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ProofOfAuthorizationRequestSubmissions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmission">ProofOfAuthorizationRequestSubmission</a>

Methods:

- <code title="post /proof_of_authorization_request_submissions">client.ProofOfAuthorizationRequestSubmissions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmissionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmissionNewParams">ProofOfAuthorizationRequestSubmissionNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmission">ProofOfAuthorizationRequestSubmission</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /proof_of_authorization_request_submissions/{proof_of_authorization_request_submission_id}">client.ProofOfAuthorizationRequestSubmissions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmissionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, proofOfAuthorizationRequestSubmissionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmission">ProofOfAuthorizationRequestSubmission</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /proof_of_authorization_request_submissions">client.ProofOfAuthorizationRequestSubmissions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmissionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmissionListParams">ProofOfAuthorizationRequestSubmissionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ProofOfAuthorizationRequestSubmission">ProofOfAuthorizationRequestSubmission</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatement">AccountStatement</a>

Methods:

- <code title="get /account_statements/{account_statement_id}">client.AccountStatements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountStatementID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatement">AccountStatement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_statements">client.AccountStatements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatementListParams">AccountStatementListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatement">AccountStatement</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#File">File</a>

Methods:

- <code title="post /files">client.Files.<a href="https://pkg.go.dev/github.com/Increase/increase-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/Increase/increase-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/Increase/increase-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#File">File</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Documents

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Document">Document</a>

Methods:

- <code title="get /documents/{document_id}">client.Documents.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /documents">client.Documents.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#DocumentListParams">DocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Document">Document</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Exports

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Export">Export</a>

Methods:

- <code title="post /exports">client.Exports.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExportService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExportNewParams">ExportNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Export">Export</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /exports/{export_id}">client.Exports.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, exportID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Export">Export</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /exports">client.Exports.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ExportListParams">ExportListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Export">Export</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Event">Event</a>

Methods:

- <code title="get /events/{event_id}">client.Events.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events">client.Events.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventListParams">EventListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Event">Event</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EventSubscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscription">EventSubscription</a>

Methods:

- <code title="post /event_subscriptions">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionNewParams">EventSubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /event_subscriptions/{event_subscription_id}">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventSubscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /event_subscriptions/{event_subscription_id}">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventSubscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionUpdateParams">EventSubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /event_subscriptions">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscriptionListParams">EventSubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#EventSubscription">EventSubscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RealTimeDecisions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecision">RealTimeDecision</a>

Methods:

- <code title="get /real_time_decisions/{real_time_decision_id}">client.RealTimeDecisions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecisionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimeDecisionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecision">RealTimeDecision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /real_time_decisions/{real_time_decision_id}/action">client.RealTimeDecisions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecisionService.Action">Action</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimeDecisionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecisionActionParams">RealTimeDecisionActionParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimeDecision">RealTimeDecision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccount">BookkeepingAccount</a>
- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingBalanceLookup">BookkeepingBalanceLookup</a>

Methods:

- <code title="post /bookkeeping_accounts">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountNewParams">BookkeepingAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccount">BookkeepingAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /bookkeeping_accounts/{bookkeeping_account_id}">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountUpdateParams">BookkeepingAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccount">BookkeepingAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_accounts">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountListParams">BookkeepingAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccount">BookkeepingAccount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_accounts/{bookkeeping_account_id}/balance">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountService.Balance">Balance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingAccountBalanceParams">BookkeepingAccountBalanceParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingBalanceLookup">BookkeepingBalanceLookup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingEntrySets

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySet">BookkeepingEntrySet</a>

Methods:

- <code title="post /bookkeeping_entry_sets">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySetNewParams">BookkeepingEntrySetNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySet">BookkeepingEntrySet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entry_sets/{bookkeeping_entry_set_id}">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingEntrySetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySet">BookkeepingEntrySet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entry_sets">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySetListParams">BookkeepingEntrySetListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntrySet">BookkeepingEntrySet</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingEntries

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntry">BookkeepingEntry</a>

Methods:

- <code title="get /bookkeeping_entries/{bookkeeping_entry_id}">client.BookkeepingEntries.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingEntryID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntry">BookkeepingEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entries">client.BookkeepingEntries.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntryListParams">BookkeepingEntryListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#BookkeepingEntry">BookkeepingEntry</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Group">Group</a>

Methods:

- <code title="get /groups/current">client.Groups.<a href="https://pkg.go.dev/github.com/Increase/increase-go#GroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OAuthConnections

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnection">OAuthConnection</a>

Methods:

- <code title="get /oauth_connections/{oauth_connection_id}">client.OAuthConnections.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, oauthConnectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnection">OAuthConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /oauth_connections">client.OAuthConnections.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnectionListParams">OAuthConnectionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthConnection">OAuthConnection</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OAuthTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthToken">OAuthToken</a>

Methods:

- <code title="post /oauth/tokens">client.OAuthTokens.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthTokenNewParams">OAuthTokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#OAuthToken">OAuthToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# IntrafiAccountEnrollments

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollment">IntrafiAccountEnrollment</a>

Methods:

- <code title="post /intrafi_account_enrollments">client.IntrafiAccountEnrollments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentNewParams">IntrafiAccountEnrollmentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollment">IntrafiAccountEnrollment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /intrafi_account_enrollments/{intrafi_account_enrollment_id}">client.IntrafiAccountEnrollments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, intrafiAccountEnrollmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollment">IntrafiAccountEnrollment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /intrafi_account_enrollments">client.IntrafiAccountEnrollments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentListParams">IntrafiAccountEnrollmentListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollment">IntrafiAccountEnrollment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /intrafi_account_enrollments/{intrafi_account_enrollment_id}/unenroll">client.IntrafiAccountEnrollments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollmentService.Unenroll">Unenroll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, intrafiAccountEnrollmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiAccountEnrollment">IntrafiAccountEnrollment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# IntrafiBalances

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiBalance">IntrafiBalance</a>

Methods:

- <code title="get /intrafi_balances/{account_id}">client.IntrafiBalances.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiBalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiBalance">IntrafiBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# IntrafiExclusions

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusion">IntrafiExclusion</a>

Methods:

- <code title="post /intrafi_exclusions">client.IntrafiExclusions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionNewParams">IntrafiExclusionNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusion">IntrafiExclusion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /intrafi_exclusions/{intrafi_exclusion_id}">client.IntrafiExclusions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, intrafiExclusionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusion">IntrafiExclusion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /intrafi_exclusions">client.IntrafiExclusions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionListParams">IntrafiExclusionListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusion">IntrafiExclusion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /intrafi_exclusions/{intrafi_exclusion_id}/archive">client.IntrafiExclusions.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusionService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, intrafiExclusionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#IntrafiExclusion">IntrafiExclusion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RealTimePaymentsRequestForPayments

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPayment">RealTimePaymentsRequestForPayment</a>

Methods:

- <code title="post /real_time_payments_request_for_payments">client.RealTimePaymentsRequestForPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPaymentNewParams">RealTimePaymentsRequestForPaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPayment">RealTimePaymentsRequestForPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_request_for_payments/{request_for_payment_id}">client.RealTimePaymentsRequestForPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestForPaymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPayment">RealTimePaymentsRequestForPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_request_for_payments">client.RealTimePaymentsRequestForPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPaymentListParams">RealTimePaymentsRequestForPaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsRequestForPayment">RealTimePaymentsRequestForPayment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Simulations

## InterestPayments

Methods:

- <code title="post /simulations/interest_payments">client.Simulations.InterestPayments.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInterestPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInterestPaymentNewParams">SimulationInterestPaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardAuthorizations

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationNewResponse">SimulationCardAuthorizationNewResponse</a>

Methods:

- <code title="post /simulations/card_authorizations">client.Simulations.CardAuthorizations.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationNewParams">SimulationCardAuthorizationNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationNewResponse">SimulationCardAuthorizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardAuthorizationExpirations

Methods:

- <code title="post /simulations/card_authorization_expirations">client.Simulations.CardAuthorizationExpirations.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationExpirationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardAuthorizationExpirationNewParams">SimulationCardAuthorizationExpirationNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardSettlements

Methods:

- <code title="post /simulations/card_settlements">client.Simulations.CardSettlements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardSettlementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardSettlementNewParams">SimulationCardSettlementNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardReversals

Methods:

- <code title="post /simulations/card_reversals">client.Simulations.CardReversals.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardReversalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardReversalNewParams">SimulationCardReversalNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardIncrements

Methods:

- <code title="post /simulations/card_increments">client.Simulations.CardIncrements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardIncrementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardIncrementNewParams">SimulationCardIncrementNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardFuelConfirmations

Methods:

- <code title="post /simulations/card_fuel_confirmations">client.Simulations.CardFuelConfirmations.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardFuelConfirmationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardFuelConfirmationNewParams">SimulationCardFuelConfirmationNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardRefunds

Methods:

- <code title="post /simulations/card_refunds">client.Simulations.CardRefunds.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardRefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardRefundNewParams">SimulationCardRefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardDisputes

Methods:

- <code title="post /simulations/card_disputes/{card_dispute_id}/action">client.Simulations.CardDisputes.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardDisputeService.Action">Action</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardDisputeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCardDisputeActionParams">SimulationCardDisputeActionParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PhysicalCards

Methods:

- <code title="post /simulations/physical_cards/{physical_card_id}/advance_shipment">client.Simulations.PhysicalCards.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationPhysicalCardService.AdvanceShipment">AdvanceShipment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationPhysicalCardAdvanceShipmentParams">SimulationPhysicalCardAdvanceShipmentParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DigitalWalletTokenRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDigitalWalletTokenRequestNewResponse">SimulationDigitalWalletTokenRequestNewResponse</a>

Methods:

- <code title="post /simulations/digital_wallet_token_requests">client.Simulations.DigitalWalletTokenRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDigitalWalletTokenRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDigitalWalletTokenRequestNewParams">SimulationDigitalWalletTokenRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDigitalWalletTokenRequestNewResponse">SimulationDigitalWalletTokenRequestNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundFundsHolds

Response Types:

- <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundFundsHoldReleaseResponse">SimulationInboundFundsHoldReleaseResponse</a>

Methods:

- <code title="post /simulations/inbound_funds_holds/{inbound_funds_hold_id}/release">client.Simulations.InboundFundsHolds.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundFundsHoldService.Release">Release</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundFundsHoldID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundFundsHoldReleaseResponse">SimulationInboundFundsHoldReleaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountTransfers

Methods:

- <code title="post /simulations/account_transfers/{account_transfer_id}/complete">client.Simulations.AccountTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationAccountTransferService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ACHTransfers

Methods:

- <code title="post /simulations/ach_transfers/{ach_transfer_id}/acknowledge">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferService.Acknowledge">Acknowledge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/create_notification_of_change">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferService.NewNotificationOfChange">NewNotificationOfChange</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferNewNotificationOfChangeParams">SimulationACHTransferNewNotificationOfChangeParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/return">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferService.Return">Return</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferReturnParams">SimulationACHTransferReturnParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/settle">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferService.Settle">Settle</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/submit">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationACHTransferService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundACHTransfers

Methods:

- <code title="post /simulations/inbound_ach_transfers">client.Simulations.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundACHTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundACHTransferNewParams">SimulationInboundACHTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## WireTransfers

Methods:

- <code title="post /simulations/wire_transfers/{wire_transfer_id}/reverse">client.Simulations.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationWireTransferService.Reverse">Reverse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/wire_transfers/{wire_transfer_id}/submit">client.Simulations.WireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationWireTransferService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundWireTransfers

Methods:

- <code title="post /simulations/inbound_wire_transfers">client.Simulations.InboundWireTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundWireTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundWireTransferNewParams">SimulationInboundWireTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireTransfer">InboundWireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundWireDrawdownRequests

Methods:

- <code title="post /simulations/inbound_wire_drawdown_requests">client.Simulations.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundWireDrawdownRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundWireDrawdownRequestNewParams">SimulationInboundWireDrawdownRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CheckTransfers

Methods:

- <code title="post /simulations/check_transfers/{check_transfer_id}/mail">client.Simulations.CheckTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCheckTransferService.Mail">Mail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundCheckDeposits

Methods:

- <code title="post /simulations/inbound_check_deposits">client.Simulations.InboundCheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundCheckDepositService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundCheckDepositNewParams">SimulationInboundCheckDepositNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundCheckDeposit">InboundCheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RealTimePaymentsTransfers

Methods:

- <code title="post /simulations/real_time_payments_transfers/{real_time_payments_transfer_id}/complete">client.Simulations.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationRealTimePaymentsTransferService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimePaymentsTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationRealTimePaymentsTransferCompleteParams">SimulationRealTimePaymentsTransferCompleteParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundRealTimePaymentsTransfers

Methods:

- <code title="post /simulations/inbound_real_time_payments_transfers">client.Simulations.InboundRealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundRealTimePaymentsTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundRealTimePaymentsTransferNewParams">SimulationInboundRealTimePaymentsTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundRealTimePaymentsTransfer">InboundRealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CheckDeposits

Methods:

- <code title="post /simulations/check_deposits/{check_deposit_id}/reject">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCheckDepositService.Reject">Reject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/check_deposits/{check_deposit_id}/return">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCheckDepositService.Return">Return</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/check_deposits/{check_deposit_id}/submit">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationCheckDepositService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundMailItems

Methods:

- <code title="post /simulations/inbound_mail_items">client.Simulations.InboundMailItems.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundMailItemService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationInboundMailItemNewParams">SimulationInboundMailItemNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#InboundMailItem">InboundMailItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Programs

Methods:

- <code title="post /simulations/programs">client.Simulations.Programs.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationProgramService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationProgramNewParams">SimulationProgramNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Program">Program</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountStatements

Methods:

- <code title="post /simulations/account_statements">client.Simulations.AccountStatements.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationAccountStatementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationAccountStatementNewParams">SimulationAccountStatementNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#AccountStatement">AccountStatement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Documents

Methods:

- <code title="post /simulations/documents">client.Simulations.Documents.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDocumentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#SimulationDocumentNewParams">SimulationDocumentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Increase/increase-go">increase</a>.<a href="https://pkg.go.dev/github.com/Increase/increase-go#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
