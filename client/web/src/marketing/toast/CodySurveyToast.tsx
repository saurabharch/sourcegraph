import { useCallback, useEffect, useState } from 'react'

import { mdiEmail } from '@mdi/js'
import classNames from 'classnames'
import { useLocation } from 'react-router-dom'

import { asError, type ErrorLike } from '@sourcegraph/common'
import type { TelemetryV2Props } from '@sourcegraph/shared/src/telemetry'
import type { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { AnchorLink, Button, H3, Icon, Modal, Text } from '@sourcegraph/wildcard'

import type { AuthenticatedUser } from '../../auth'
import { getReturnTo } from '../../auth/SignInSignUpCommon'
import { CodyProRoutes } from '../../cody/codyProRoutes'
import { resendVerificationEmail } from '../../user/settings/emails/UserEmail'

import styles from './CodySurveyToast.module.scss'

const CodyVerifyEmailToast: React.FC<
    { onNext: () => void; authenticatedUser: AuthenticatedUser } & TelemetryProps & TelemetryV2Props
> = ({ onNext, authenticatedUser, telemetryService, telemetryRecorder }) => {
    const [sending, setSending] = useState(false)
    const [resentEmailTo, setResentEmailTo] = useState<string | null>(null)
    const [resendEmailError, setResendEmailError] = useState<ErrorLike | null>(null)
    const resend = useCallback(async () => {
        const email = (authenticatedUser.emails || []).find(({ verified }) => !verified)?.email
        if (email) {
            setSending(true)
            await resendVerificationEmail(authenticatedUser.id, email, telemetryRecorder, {
                onSuccess: () => {
                    setResentEmailTo(email)
                    setResendEmailError(null)
                    setSending(false)
                },
                onError: (errors: ErrorLike) => {
                    setResendEmailError(asError(errors))
                    setResentEmailTo(null)
                    setSending(false)
                },
            })
        }
    }, [authenticatedUser, telemetryRecorder])

    useEffect(() => {
        telemetryService.log('VerifyEmailToastViewed')
        telemetryRecorder.recordEvent('codySurvey.veryEmailToast', 'view')
    }, [telemetryService, telemetryRecorder])

    return (
        <Modal
            className={styles.codySurveyToastModal}
            position="center"
            aria-label="Welcome message"
            containerClassName={styles.modalOverlay}
        >
            <H3 className="mb-4">
                <Icon svgPath={mdiEmail} className={classNames('mr-2', styles.emailIcon)} aria-hidden={true} />
                Verify your email address
            </H3>
            <Text>To use Cody, our AI Assistant, you'll need to verify your email address.</Text>
            <Text className="d-flex align-items-center">
                <span className="mr-1">Didn't get an email?</span>
                {sending ? (
                    <span>Sending...</span>
                ) : (
                    <>
                        <span>Click to </span>
                        <Button variant="link" className={classNames('p-0 ml-1', styles.resendButton)} onClick={resend}>
                            resend
                        </Button>
                        .
                    </>
                )}
            </Text>
            {resentEmailTo && (
                <Text>
                    Sent verification email to <strong>{resentEmailTo}</strong>.
                </Text>
            )}
            {resendEmailError && <Text>{resendEmailError.message}.</Text>}
            <div className="d-flex justify-content-end mt-4">
                <AnchorLink className="mr-3 mt-auto mb-auto" to="/-/sign-out">
                    Sign out
                </AnchorLink>
                <Button className={styles.codySurveyToastModalButton} variant="primary" onClick={onNext}>
                    Next
                </Button>
            </div>
        </Modal>
    )
}

export const CodySurveyToast: React.FC<
    {
        authenticatedUser: AuthenticatedUser
    } & TelemetryProps &
        TelemetryV2Props
> = ({ authenticatedUser, telemetryService, telemetryRecorder }) => {
    const [showVerifyEmail, setShowVerifyEmail] = useState(!authenticatedUser.hasVerifiedEmail)

    const location = useLocation()

    const dismissVerifyEmail = useCallback(() => {
        telemetryService.log('VerifyEmailToastDismissed')
        telemetryRecorder.recordEvent('codySurvey.verifyEmailToast', 'dismissed')
        setShowVerifyEmail(false)
    }, [telemetryService, telemetryRecorder])

    useEffect(() => {
        telemetryService.log('CustomerQualificationSurveyExperiment001Enrolled')
        telemetryRecorder.recordEvent('experiment', 'enroll', { metadata: { experimentId: 1 } })
    }, [telemetryService, telemetryRecorder])

    if (!showVerifyEmail) {
        // Redirects once user submits the post-sign-up form
        const returnTo = getReturnTo(location, CodyProRoutes.Manage)
        window.location.replace(returnTo)
        return null
    }

    return (
        <CodyVerifyEmailToast
            onNext={dismissVerifyEmail}
            authenticatedUser={authenticatedUser}
            telemetryService={telemetryService}
            telemetryRecorder={telemetryRecorder}
        />
    )
}
