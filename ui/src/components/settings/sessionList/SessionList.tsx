import UAParser from 'ua-parser-js';
import styles from './SessionList.module.scss';
import { useTranslation } from 'react-i18next';
import { Close } from '@/assets';
import { useLoading } from '@/contexts/LoadingContext';
import { Session } from '@/clients/api/api.types';
import Headline from '@/components/common/headline/Headline';
import Button from '@/components/common/button/Button';
import { TestIds } from '@/__test__/ids';

type Props = {
  sessions: Session[];
  onRevoke: (sessionId: number) => void;
};

const unixNow = () => Math.floor(Date.now() / 1000);

const SessionList = (props: Props) => {
  const { t } = useTranslation();
  const { isLoading } = useLoading();

  const parseUserAgent = (rawUserAgent?: string) => {
    if (!rawUserAgent) {
      return t('unknown');
    }

    const parser = new UAParser(rawUserAgent);
    const { browser, os } = parser.getResult();

    if (browser.name && browser.major && os.name && os.version) {
      return t('user_agent_full', {
        browser: browser,
        os: os,
      });
    } else if (browser.name && browser.major) {
      return t('user_agent_browser', {
        browser: browser,
      });
    } else if (browser.name) {
      return browser.name;
    } else if (os.name) {
      return os.name;
    }

    return rawUserAgent;
  };

  const formatDate = (date: number) => {
    const now = unixNow();
    const nowThreshold = 60;

    if (now - nowThreshold > date) {
      return t('date', { date });
    }

    return t('now');
  };

  return (
    <div className={styles.SessionList}>
      <Headline>{t('sessions')}</Headline>
      <div className={styles.tableContainer}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>{t('last_seen')}</th>
              <th className={styles.created}>{t('created')}</th>
              <th className={styles.ip}>{t('ip_address')}</th>
              <th>{t('browser_and_os')}</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {props.sessions.map((session) => {
              return (
                <tr key={session.id}>
                  <td>{formatDate(session.lastUsedAt)}</td>
                  <td className={styles.created}>
                    {formatDate(session.createdAt)}
                  </td>
                  <td className={styles.ip}>{session.clientIp}</td>
                  <td>{parseUserAgent(session.userAgent)}</td>
                  <td>
                    <Button
                      icon={<Close className={styles.sessionDeleteButton} />}
                      style="red"
                      onClick={() => props.onRevoke(session.id)}
                      disabled={isLoading}
                      testID={TestIds.REVOKE_SESSION_BUTTON}
                    />
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default SessionList;
