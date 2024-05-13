import { useEffect } from 'react'

import { Elements } from '@stripe/react-stripe-js'
// NOTE: A side effect of loading this library will update the DOM and
// fetch stripe.js. This is a subtle detail but means that the Stripe
// functionality won't be loaded until this actual module does, via
// the lazily loaded router module.
import * as stripeJs from '@stripe/stripe-js'
import classNames from 'classnames'
import { Navigate } from 'react-router-dom'

import { useQuery } from '@sourcegraph/http-client'
import type { TelemetryV2Props } from '@sourcegraph/shared/src/telemetry'
import { PageHeader, LoadingSpinner } from '@sourcegraph/wildcard'

import type { AuthenticatedUser } from '../../../../auth'
import { withAuthenticatedUser } from '../../../../auth/withAuthenticatedUser'
import { Page } from '../../../../components/Page'
import { PageTitle } from '../../../../components/PageTitle'
import { type UserCodyPlanResult, type UserCodyPlanVariables } from '../../../../graphql-operations'
import { USER_CODY_PLAN } from '../../../subscription/queries'

import { testInvoices, testSubscription } from './testData'
import { InvoiceHistoryCardComponent } from './InvoiceHistoryCardComponent'
import { SubscriptionDetailsCardComponent } from './SubscriptionDetailsCardComponent'

// TODO: Figure out how to unify this with the duplication in NewCodyProSubscriptionPage.
// Can we have a singleton for loading Stripe or something?
//
// NOTE: Call loadStripe outside a component’s render to avoid recreating the object.
// We do it here, meaning that "stripe.js" will get loaded lazily, when the user
// routes to this page.
const publishableKey = window.context.frontendCodyProConfig?.stripePublishableKey
const stripePromise = stripeJs.loadStripe(publishableKey || '')

interface ManageCodyProSubscriptionPageProps extends TelemetryV2Props {
    authenticatedUser: AuthenticatedUser
}

const AuthenticatedManageCodyProSubscriptionPage: React.FunctionComponent<ManageCodyProSubscriptionPageProps> = ({
    authenticatedUser,
    telemetryRecorder,
}) => {
    console.log('AuthenticatedMangeCodyProSubscriptionPage')

    useEffect(() => {
        telemetryRecorder.recordEvent('cody.manage-subscription', 'view')
    }, [telemetryRecorder])

    // This page only applies to users who have a Cody Pro subscription to manage.
    // Otherwise, direct them to the ./new page to sign up.
    const { data, error: dataLoadError } = useQuery<UserCodyPlanResult, UserCodyPlanVariables>(USER_CODY_PLAN, {})
    if (dataLoadError) {
        throw dataLoadError
    }
    const subscriptionData = data?.currentUser?.codySubscription
    if (!subscriptionData) {
        return <LoadingSpinner></LoadingSpinner>
    }
    if (subscriptionData.plan !== 'PRO') {
        return <Navigate to="/cody/manage/subscription/new" replace={true} />
    }

    // TODO: Reconcile the Sourcegraph.com GraphQL and the SSC backend's REST API.
    // We may use either or even both to power this page. But as a first step, we
    // just have stubbed out data to get the UI situated.
    const invoices = testInvoices
    const subscription = testSubscription

    // If both the team and subscription details are defined, then the user is a team admin
    // and has full access to modify subscription data as needed.
    const appearance: stripeJs.Appearance = {
        theme: 'stripe',
        variables: {
            colorPrimary: '#00b4d9',
        },
    }
    return (
        <>
            <Page className={classNames('d-flex flex-column')}>
                <PageTitle title="Manage Subscription" />
                <PageHeader className="mb-4 mt-4">
                    <PageHeader.Heading as="h2" styleAs="h1">
                        <div className="d-inline-flex align-items-center">Manage Subscription</div>
                    </PageHeader.Heading>
                </PageHeader>

                <Elements stripe={stripePromise} options={{ appearance }}></Elements>
                <p>
                    <a href="/cody/manage" className="flex items-center gap-2 mb-6">
                        Back to Cody Dashboard
                    </a>
                </p>

                <div className="pb-4">
                    <SubscriptionDetailsCardComponent subscription={subscription}></SubscriptionDetailsCardComponent>
                </div>

                <InvoiceHistoryCardComponent invoices={invoices}></InvoiceHistoryCardComponent>
            </Page>
        </>
    )
}

export const ManageCodyProSubscriptionPage = withAuthenticatedUser(AuthenticatedManageCodyProSubscriptionPage)
