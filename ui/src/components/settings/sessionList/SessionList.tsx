import Button from '@/components/common/button/Button';
import { Close } from '@/assets';
import Headline from '@/components/common/headline/Headline';
import { Session } from '@/clients/api/api.types';
import Table from '@/components/common/table/Table';
import TableBody from '@/components/common/table/TableBody';
import TableCell from '@/components/common/table/TableCell';
import TableHead from '@/components/common/table/TableHead';
import TableHeadCell from '@/components/common/table/TableHeadCell';
import TableRow from '@/components/common/table/TableRow';
import { TestIds } from '@/__test__/ids';
import UAParser from 'ua-parser-js';
import styles from './SessionList.module.scss';
import { useLoading } from '@/contexts/LoadingContext';
import { useTranslation } from 'react-i18next';

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

  const renderTable = () => {
    if (props.sessions.length === 0) {
      return <div>{t('no_sessions')}</div>;
    }

    return (
      <div className={styles.tableContainer}>
        <Table className={styles.table}>
          <TableHead>
            <TableHeadCell>{t('last_seen')}</TableHeadCell>
            <TableHeadCell className={styles.created}>
              {t('created')}
            </TableHeadCell>
            <TableHeadCell className={styles.ip}>
              {t('ip_address')}
            </TableHeadCell>
            <TableHeadCell>{t('browser_and_os')}</TableHeadCell>
            <TableHeadCell />
          </TableHead>
          <TableBody>
            {props.sessions.map((session) => {
              return (
                <TableRow key={session.id}>
                  <TableCell>{formatDate(session.lastUsedAt)}</TableCell>
                  <TableCell className={styles.created}>
                    {formatDate(session.createdAt)}
                  </TableCell>
                  <TableCell className={styles.ip}>
                    {session.clientIp}
                  </TableCell>
                  <TableCell>{parseUserAgent(session.userAgent)}</TableCell>
                  <TableCell>
                    <Button
                      icon={<Close className={styles.sessionDeleteButton} />}
                      style="red"
                      onClick={() => props.onRevoke(session.id)}
                      disabled={isLoading}
                      testID={TestIds.REVOKE_SESSION_BUTTON}
                    />
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </div>
    );
  };

  return (
    <div className={styles.SessionList}>
      <Headline>{t('sessions')}</Headline>
      {renderTable()}
    </div>
  );
};

export default SessionList;
