import { Container, Link, H2 } from '@sourcegraph/wildcard'

import { Invoice } from './apitypes'
import { humanizeDate, usdCentsToHumanString } from './util'

export interface Invoice./testDatadProps {
    invoices: Invoice[]
}

export const InvoiceHistoryCardComponent: React.FunctionComponent<InvoiceHistoryCardProps> = ({ invoices }) => {
    return (
        <Container>
            <H2>Invoice History</H2>
            <hr />
            <div className="d-flex justify-content-between align-items-center">
                {invoices.length ? (
                    <ul>
                        {invoices.map((invoice, index) => {
                            return (
                                <div key={index}>
                                    <span>
                                        📄 {invoice.periodEnd ? humanizeDate(invoice.periodEnd as string) : '(no date)'}
                                    </span>
                                    <span>{usdCentsToHumanString(invoice.amountDue)}</span>
                                    <span>
                                        {invoice.status.charAt(0).toUpperCase() + invoice.status.slice(1).toLowerCase()}
                                    </span>
                                    {invoice.hostedInvoiceUrl ? (
                                        <Link to={invoice.hostedInvoiceUrl} target="_blank">
                                            Get invoice
                                        </Link>
                                    ) : (
                                        '-'
                                    )}
                                </div>
                            )
                        })}
                    </ul>
                ) : (
                    <p>You have no invoices.</p>
                )}
            </div>
        </Container>
    )
}
