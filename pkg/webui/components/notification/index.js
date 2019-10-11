// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import classnames from 'classnames'

import Button from '../button'
import Message from '../../lib/components/message'
import Icon from '../icon'
import Details from './details'
import PropTypes from '../../lib/prop-types'
import style from './notification.styl'

const Notification = function({
  className,
  title,
  error,
  warning,
  info,
  small,
  message,
  success,
  messageValues = {},
  action,
  actionMessage,
  buttonIcon,
  children,
  details,
}) {
  const classname = classnames(style.notification, className, {
    [style.error]: error,
    [style.warning]: warning,
    [style.info]: info,
    [style.small]: small,
    [style.success]: success,
    [style.withDetails]: Boolean(details),
  })

  let icon = 'info'
  if (error) {
    icon = 'error'
  } else if (warning) {
    icon = 'warning'
  }

  const content = message || error || warning || info || success

  return (
    <div className={classname}>
      <div className={style.container}>
        <Icon className={style.icon} icon={icon} large={!small} />
        <div className={style.content}>
          {title && <Message className={style.title} content={title} component="h4" />}
          <div>
            {content && <Message content={content} values={messageValues} />}
            {children}
            {action && (
              <Button naked secondary icon={buttonIcon} onClick={action} message={actionMessage} />
            )}
          </div>
        </div>
      </div>
      {details && <Details className={style.details} details={details} />}
    </div>
  )
}

Notification.propTypes = {
  action: PropTypes.func,
  actionMessage: PropTypes.message,
  buttonIcon: PropTypes.string,
  details: PropTypes.mixed,
  children: PropTypes.node,
  error: PropTypes.error,
  info: PropTypes.message,
  message: PropTypes.message,
  messageValues: PropTypes.object,
  small: PropTypes.bool,
  success: PropTypes.message,
  title: PropTypes.message,
  warning: PropTypes.message,
}

Notification.defaultProps = {
  action: undefined,
  actionMessage: undefined,
  buttonIcon: '',
  details: undefined,
}

export default Notification
