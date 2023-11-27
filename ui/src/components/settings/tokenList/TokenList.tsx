import { Close } from '@/assets';
import { Token } from '@/clients/api/api.types';
import { useLoading } from '@/contexts/LoadingContext';
import styles from './TokenList.module.scss';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import Headline from '@/components/common/headline/Headline';
import Button from '@/components/common/button/Button';

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
        <table className={styles.table}>
          <thead>
            <tr>
              <th>{t('created')}</th>
              <th>{t('name')}</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {props.tokens.map((token) => {
              return (
                <tr key={token.id}>
                  <td>{formatDate(token.createdAt)}</td>
                  <td>{token.name}</td>
                  <td>
                    <Button
                      icon={<Close className={styles.deleteButton} />}
                      style="red"
                      onClick={() => props.onRevoke(token.id)}
                      disabled={isLoading}
                    />
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
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
