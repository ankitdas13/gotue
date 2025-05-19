# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListPaymentsResponse">OrderListPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListPaymentsParams">OrderListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#OrderListPaymentsResponse">OrderListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#AcquirerData">AcquirerData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardDetails

Methods:

- <code title="get /payments/{id}/card">client.Payments.CardDetails.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentCardDetailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundList">RefundList</a>

Methods:

- <code title="post /payments/{id}/refund">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundNewParams">PaymentRefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundGetParams">PaymentRefundGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentRefundListParams">PaymentRefundListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#RefundUpdateNotesParams">RefundUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go">gotue</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gotue-go#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
