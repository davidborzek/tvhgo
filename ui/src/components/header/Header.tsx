import { BurgerMenuIcon, SearchIcon, TvhgoHorizontalLogo } from '@/assets';
import { useMatch, useNavigate, useSearchParams } from 'react-router-dom';

import Input from '../common/input/Input';
import styles from './Header.module.scss';
import { useFormik } from 'formik';
import { useTranslation } from 'react-i18next';

type Props = {
  onToggle: () => void;
};

function Header({ onToggle }: Props) {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const isSearch = useMatch(`/search`);
  const [searchParams] = useSearchParams();

  const searchForm = useFormik({
    initialValues: {
      query: isSearch ? searchParams.get('q') || '' : '',
    },
    onSubmit: ({ query }) => navigate(`/search?q=${encodeURIComponent(query)}`),
    enableReinitialize: true,
  });

  const renderSearch = () => {
    return (
      <form onSubmit={searchForm.handleSubmit}>
        <Input
          name="query"
          icon={<SearchIcon />}
          placeholder={t('search')}
          value={searchForm.values.query}
          onChange={searchForm.handleChange}
          onBlur={searchForm.handleBlur}
        />
      </form>
    );
  };

  return (
    <div className={styles.root}>
      <div className={styles.bar}>
        <div className={styles.left}>
          <BurgerMenuIcon
            className={styles.menuIcon}
            onClick={() => onToggle()}
          />
          <TvhgoHorizontalLogo className={styles.logo} />
        </div>
        <div className={styles.right}>{renderSearch()}</div>
      </div>
    </div>
  );
}

export default Header;
