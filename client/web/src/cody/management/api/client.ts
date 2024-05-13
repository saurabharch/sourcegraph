import * as types from './types'

// Client provides the capabilities exposed from the backend Cody Pro REST API.
//
// Exposed as an interface to mock as necessary for testing.
export interface Client {
    // Subscriptions
    getCurrentTeamSubscription(): Promise<types.Subscription>
    getCurrentTeamSubscriptionSummary(): Promise<types.SubscriptionSummary>
    updateCurrentTeamSubscription(_: types.UpdateSubscriptionRequest): Promise<types.Subscription>
    getCurrentTeamSubscriptionInvoices(): Promise<types.GetSubscriptionInvoicesResponse>
    reactivateCurrentTeamSubscription(_: types.ReactivateSubscriptionRequest): Promise<types.Subscription>

    // Team Members
    getCurrentTeamMembers(): Promise<types.ListTeamInvitesResponse>
    updateTeamMembers(_: types.UpdateTeamMembersRequest): Promise<void>

    // Team Invites
    getCurrentTeamInvites(): Promise<types.ListTeamInvitesResponse>
    createTeamInvite(_: types.CreateTeamInviteRequest): Promise<types.TeamInvite>
    cancelCurrentTeamInvite(inviteId: string): Promise<void>
    getTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite>
    acceptTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite>
    cancelTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite>

    // Stripe Checkout
    createStripeCheckoutSession(_: types.CreateCheckoutSessionRequest): Promise<types.CreateCheckoutSessionResponse>
    getCheckoutSession(sessionId: string): Promise<types.GetCheckoutSessionResponse>
}

export class CodyProBackendClient implements Client {
    // e.g. "https://sourcegraph.com"
    private origin: string

    constructor() {
        this.origin = window.location.origin
    }

    private async get<Resp>(urlSuffix: string): Promise<Resp> {
        return this.call<undefined, Resp>('GET', urlSuffix, undefined /* body */)
    }

    private async patch<Req, Resp>(urlSuffix: string, body: Req): Promise<Resp> {
        return this.call('PATCH', urlSuffix, body)
    }

    private async post<Req, Resp>(urlSuffix: string, body: Req): Promise<Resp> {
        return this.call('POST', urlSuffix, body)
    }

    private async call<Req, Resp>(method: string, urlSuffix: string, body: Req): Promise<Resp> {
        const response = await fetch(`${this.origin}/.api/ssc/proxy${urlSuffix}`, {
            // Pass along the "sgs" session cookie to identify the caller.
            credentials: 'same-origin',
            method,
            body: JSON.stringify(body),
        })

        const responseBody = await response.text()
        if (response.status >= 200 && response.status <= 299) {
            const typedResp = JSON.parse(responseBody) as Resp
            return typedResp
        }

        // NOTE: We'll want to update the REST API on the SSC backend to return
        // 4xx responses in a recognizable JSON format. That way callers will
        // be able to catch and handle common errors.
        throw Error(`Unexpected status calling backend API: ${response.status}`)
    }

    // Subscriptions
    getCurrentTeamSubscription(): Promise<types.Subscription> {
        return this.get('/team/current/subscription')
    }

    getCurrentTeamSubscriptionSummary(): Promise<types.SubscriptionSummary> {
        return this.get('/team/current/subscription/summary')
    }

    updateCurrentTeamSubscription(req: types.UpdateSubscriptionRequest): Promise<types.Subscription> {
        return this.patch('/team/current/subscription', req)
    }

    getCurrentTeamSubscriptionInvoices(): Promise<types.GetSubscriptionInvoicesResponse> {
        return this.get('/team/current/subscription/invoices')
    }

    reactivateCurrentTeamSubscription(req: types.ReactivateSubscriptionRequest): Promise<types.Subscription> {
        return this.post('/team/current/subscription/summary', req)
    }

    // Team Members
    getCurrentTeamMembers(): Promise<types.ListTeamInvitesResponse> {
        return this.get('/team/current/members')
    }

    updateTeamMembers(req: types.UpdateTeamMembersRequest): Promise<void> {
        return this.patch('team/current/members', req)
    }

    // Team Invites
    getCurrentTeamInvites(): Promise<types.ListTeamInvitesResponse> {
        return this.get('/team/current/invites')
    }

    createTeamInvite(req: types.CreateTeamInviteRequest): Promise<types.TeamInvite> {
        return this.post('/team/current/invites', req)
    }

    cancelCurrentTeamInvite(inviteId: string): Promise<void> {
        return this.post(`/team/current/${inviteId}/cancel`, {})
    }

    getTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite> {
        return this.get(`/team/${teamId}/invites/${inviteId}`)
    }

    acceptTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite> {
        return this.post(`/team/${teamId}/invites/${inviteId}/accept`, {})
    }

    cancelTeamInvite(teamId: string, inviteId: string): Promise<types.TeamInvite> {
        return this.post(`/team/${teamId}/invites/${inviteId}/cancel`, {})
    }

    // Stripe Checkout
    createStripeCheckoutSession(req: types.CreateCheckoutSessionRequest): Promise<types.CreateCheckoutSessionResponse> {
        return this.post('/checkout/session', req)
    }

    getCheckoutSession(sessionId: string): Promise<types.GetCheckoutSessionResponse> {
        return this.get(`/checkout/session?session_id=${sessionId}`)
    }
}
