import UAParser from 'ua-parser-js';
import { Session } from '../../clients/api/api.types';
import styles from './SessionList.module.scss';
import { useTranslation } from 'react-i18next';
import Button from '../Button/Button';

type Props = {
  sessions: Session[];
  onRevoke: (sessionId: number) => void;
};

const unixNow = () => Math.floor(Date.now() / 1000);

const SessionList = (props: Props) => {
  const { t } = useTranslation();

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
      <h2>{t('sessions')}</h2>
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
                      label={t('delete')}
                      style="red"
                      onClick={() => props.onRevoke(session.id)}
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
