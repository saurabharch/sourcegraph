import { Container, Text } from '@sourcegraph/wildcard'

import { Address, Subscription } from '../../api/types'
import { humanizeBillingInterval, usdCentsToHumanString } from './util'

export interface SubscriptionDetailsCardComponentProps {
    subscription: Subscription
}

export const SubscriptionDetailsCardComponent: React.FunctionComponent<SubscriptionDetailsCardComponentProps> = ({
    subscription,
}) => {
    // These call outs are inform the type system not to complain that these fields may be undefined.
    if (!subscription.nextInvoice) {
        return <Text>TODO: Support alternate states, such as canceled subscriptions.</Text>
    }

    // The Customer account data isn't attached to the Subscription object, which is probably a bug
    // on the backend. We need to expose things like payment method, billing address, etc.
    const customerName = 'Custoner Name Goes Here'
    const paymentMethod = {
        expMonth: 6,
        expYear: 30,
        last4: '4242',
    }
    const address: Address = {
        line1: 'Evergreen Terrance',
        line2: '',
        city: 'Springfield',
        state: 'IL',
        country: 'US',
        postalCode: '98052',
    }
    return (
        <Container>
            {/* Top section, with the amount and renewal date. */}
            <div>
                <Text>
                    {usdCentsToHumanString(subscription.nextInvoice.newPrice)}/
                    {humanizeBillingInterval(subscription.billingInterval)}
                </Text>
            </div>

            <hr />

            {/* Payment method. */}
            <div>
                <Text>
                    Expiration {paymentMethod.expMonth}/{paymentMethod.expYear}
                </Text>
                <Text>x{paymentMethod.last4}</Text>
            </div>

            {/* Billing Address. */}
            <div>
                <Text>Full Name: {customerName}</Text>
                <Text>Country or Region: {address.country}</Text>
                <Text>Address Line 1: {address.line1}</Text>
                <Text>Address Line 2: {address.line2 || '-'}</Text>
                <Text>City: {address.city}</Text>
                <Text>State: {address.state}</Text>
                <Text>Postal Code: {address.postalCode}</Text>
            </div>
        </Container>
    )
}
