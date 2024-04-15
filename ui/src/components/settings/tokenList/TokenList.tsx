import Button from '@/components/common/button/Button';
import { Close } from '@/assets';
import Headline from '@/components/common/headline/Headline';
import Table from '@/components/common/table/Table';
import TableBody from '@/components/common/table/TableBody';
import TableCell from '@/components/common/table/TableCell';
import TableHead from '@/components/common/table/TableHead';
import TableHeadCell from '@/components/common/table/TableHeadCell';
import TableRow from '@/components/common/table/TableRow';
import { TestIds } from '@/__test__/ids';
import { Token } from '@/clients/api/api.types';
import styles from './TokenList.module.scss';
import { useLoading } from '@/contexts/LoadingContext';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

type Props = {
  tokens: Token[];
  onRevoke: (sessionId: number) => void;
};

const unixNow = () => Math.floor(Date.now() / 1000);

const TokenList = (props: Props) => {
  const { t } = useTranslation();
  const { isLoading } = useLoading();
  const navigate = useNavigate();

  const formatDate = (date: number) => {
    const now = unixNow();
    const nowThreshold = 60;

    if (now - nowThreshold > date) {
      return t('date', { date });
    }

    return t('now');
  };

  const renderTable = () => {
    return (
      <div className={styles.tableContainer}>
        <Table className={styles.table}>
          <TableHead>
            <TableHeadCell>{t('created')}</TableHeadCell>
            <TableHeadCell>{t('name')}</TableHeadCell>
            <TableHeadCell></TableHeadCell>
          </TableHead>
          <TableBody>
            {props.tokens.map((token) => {
              return (
                <TableRow key={token.id}>
                  <TableCell>{formatDate(token.createdAt)}</TableCell>
                  <TableCell>{token.name}</TableCell>
                  <TableCell>
                    <Button
                      icon={<Close className={styles.deleteButton} />}
                      style="red"
                      onClick={() => props.onRevoke(token.id)}
                      disabled={isLoading}
                      testID={TestIds.REVOKE_TOKEN_BUTTON}
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
    <div className={styles.TokenList}>
      <Headline>{t('api_tokens')}</Headline>

      {props.tokens.length > 0 ? renderTable() : <></>}

      <div>
        <Button
          label={t('new_token')}
          onClick={() => navigate('/settings/security/tokens/create')}
        />
      </div>
    </div>
  );
};

export default TokenList;
