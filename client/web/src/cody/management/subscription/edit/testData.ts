import { Invoice, Subscription } from '../../api/types'

// Stub data to populate the UI until we get the REST API wired up.
export const testInvoices: Invoice[] = [
    {
        date: new Date('2024-03-14T22:11:32Z'),
        amountDue: 15000,
        amountPaid: 0,
        status: 'open',
        periodStart: new Date('2024-03-14T22:11:32Z'),
        periodEnd: new Date('2024-04-14T22:11:32Z'),
        pdfUrl: 'https://pay.stripe.com/invoice/...',
    },
]

export const testSubscription: Subscription = {
    // BUG? If the Subscription type doesn't contain the customer contact info, then
    // where is this supposed to come from? Seems like something is missing on the REST
    // API side.

    // primaryEmail: 'email@example.com',
    // name: 'Backend Infrastructure Team',
    // address: {
    //     line1: '742 Evergreen Terrace',
    //     line2: '',
    //     city: 'Springfield',
    //     state: 'IL',
    //     postalCode: '62629',
    //     country: 'US',
    // },
    // paymentMethod: {
    //     expMonth: 6,
    //     expYear: 30,
    //     last4: '4242',
    // }
    discountInfo: {
        description: 'Test Data Discount',
        expiresAt: undefined,
    },
    subscriptionStatus: 'active',
    createdAt: new Date('2024-02-14T22:05:02Z'),
    endedAt: undefined,
    maxSeats: 100,
    billingInterval: 'monthly',
    cancelAtPeriodEnd: false,
    currentPeriodStart: new Date('2024-04-14T22:11:32Z'),
    currentPeriodEnd: new Date('2024-05-14T22:11:32Z'),
    nextInvoice: {
        newPrice: 15000,
        dueNow: 15000,
        dueDate: new Date('2024-05-14T22:11:32Z'),
    },
}
