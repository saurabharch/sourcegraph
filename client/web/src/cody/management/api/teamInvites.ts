export type TeamRole = 'none' | 'member' | 'admin'

export type TeamInviteStatus = 'sent' | 'errored' | 'accepted' | 'canceled'

export interface TeamInvite {
    id: string
    email: string
    role: TeamRole

    status: TeamInviteStatus

    // If the invite is in an errored state, this will contain a human
    // friendly description of the problem. (e.g. "unable to accept because
    // no seats are available")
    error?: string

    sentAt: Date
    acceptedAt?: Date
}

export interface CreateTeamInviteRequest {
    email: string
    role: TeamRole
}

export interface ListTeamInvitesResponse {
    invites: TeamInvite[]
    continuationToken?: string
}
