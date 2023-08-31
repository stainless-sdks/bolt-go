# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Profile">Profile</a>

Methods:

- <code title="get /account">client.Accounts.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountGetParams">AccountGetParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account/exists">client.Accounts.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountService.Exists">Exists</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountExistsParams">AccountExistsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddress">AccountAddress</a>

Methods:

- <code title="post /account/addresses">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressNewParams">AccountAddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddress">AccountAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /account/addresses/{id}">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressUpdateParams">AccountAddressUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddress">AccountAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /account/addresses/{id}">client.Accounts.Addresses.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountAddressDeleteParams">AccountAddressDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## PaymentMethods

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentMethod">PaymentMethod</a>

Methods:

- <code title="post /account/payment-methods">client.Accounts.PaymentMethods.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountPaymentMethodService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountPaymentMethodNewParams">AccountPaymentMethodNewParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentMethod">PaymentMethod</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /account/payment-methods/{id}">client.Accounts.PaymentMethods.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountPaymentMethodService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountPaymentMethodDeleteParams">AccountPaymentMethodDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentNewResponse">PaymentNewResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Guest

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#GuestPaymentNewResponse">GuestPaymentNewResponse</a>

Methods:

- <code title="post /guest/payments">client.Guest.Payments.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#GuestPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#GuestPaymentNewParams">GuestPaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#GuestPaymentNewResponse">GuestPaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Merchant

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#CallbackURLs">CallbackURLs</a>
- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Identifiers">Identifiers</a>

## Callbacks

Methods:

- <code title="patch /merchant/callbacks">client.Merchant.Callbacks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#MerchantCallbackService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#MerchantCallbackUpdateParams">MerchantCallbackUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#CallbackURLs">CallbackURLs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /merchant/callbacks">client.Merchant.Callbacks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#MerchantCallbackService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#MerchantCallbackListParams">MerchantCallbackListParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#CallbackURLs">CallbackURLs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Identifiers

Methods:

- <code title="get /merchant/identifiers">client.Merchant.Identifiers.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#MerchantIdentifierService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Identifiers">Identifiers</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookListResponse">WebhookListResponse</a>

Methods:

- <code title="get /webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookUpdateParams">WebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookListParams">WebhookListParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Testing

Response Types:

- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountTest">AccountTest</a>
- <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#CreditCard">CreditCard</a>

## Accounts

Methods:

- <code title="post /testing/accounts">client.Testing.Accounts.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#TestingAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#TestingAccountNewParams">TestingAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#AccountTest">AccountTest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CreditCards

Methods:

- <code title="get /testing/credit-cards">client.Testing.CreditCards.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#TestingCreditCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#CreditCard">CreditCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipments

Methods:

- <code title="post /testing/shipments">client.Testing.Shipments.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#TestingShipmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bolt/bolt-go">bolt</a>.<a href="https://pkg.go.dev/github.com/bolt/bolt-go#TestingShipmentNewParams">TestingShipmentNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
