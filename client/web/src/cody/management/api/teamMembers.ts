import { TeamRole } from './teamInvites'

export interface TeamMember {
    // SAMS Account External ID. We should not display this anywhere in the
    // UI, favoring the Display Name instead. But this is the stable, unique
    // identifier for the user.
    accountId: string
    displayName: string
    avatarUrl: string

    role: TeamRole
}

export interface TeamMemberRef {
    accountId: string
    teamRole: TeamRole
}

export interface ListTeamMembersResponse {
    members: TeamMember[]
    continuationToken?: string
}

export interface UpdateTeamMembersRequest {
    addMember?: TeamMemberRef
    removeMember?: TeamMemberRef
    updateMemberRole?: TeamMemberRef
}
