import H from 'history'
import React, { useEffect, useState } from 'react'
import { from, of, Subscription } from 'rxjs'
import { catchError, startWith, switchMap } from 'rxjs/operators'
import { ExtensionsControllerProps } from '../../../../../../shared/src/extensions/controller'
import { PlatformContextProps } from '../../../../../../shared/src/platform/context'
import { asError, ErrorLike } from '../../../../../../shared/src/util/errors'
import { DiagnosticsList } from '../../../tasks/list/DiagnosticsList'
import { DiagnosticInfo, toDiagnosticInfos } from '../../../threads/detail/backend'
import { StatusAreaContext } from '../StatusArea'

interface Props extends Pick<StatusAreaContext, 'status'>, ExtensionsControllerProps, PlatformContextProps {
    className?: string
    history: H.History
    location: H.Location
    isLightTheme: boolean
}

const LOADING: 'loading' = 'loading'

/**
 * The status checks page.
 */
export const StatusChecksPage: React.FunctionComponent<Props> = ({ status, className = '', ...props }) => {
    const [diagnosticsOrError, setDiagnosticsOrError] = useState<typeof LOADING | DiagnosticInfo[] | ErrorLike>(LOADING)
    useEffect(() => {
        const subscriptions = new Subscription()
        subscriptions.add(
            from(status.status.diagnostics || of([]))
                .pipe(
                    switchMap(diagEntries => toDiagnosticInfos(diagEntries)),
                    catchError(err => [asError(err)]),
                    startWith(LOADING)
                )
                .subscribe(setDiagnosticsOrError)
        )
        return () => subscriptions.unsubscribe()
    }, [status.status.diagnostics])

    return (
        <div className={`status-checks-page ${className}`}>
            <DiagnosticsList {...props} diagnosticsOrError={diagnosticsOrError} itemClassName="container-fluid" />
        </div>
    )
}