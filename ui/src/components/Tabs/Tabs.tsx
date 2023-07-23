import styles from './Tabs.module.scss';
import { c } from '../../utils/classNames';

export type Tab = {
  label: string;
  active: boolean;
};

type TabItemProps = {
  tab: Tab;
  onClick: () => void;
};

const TabItem = (props: TabItemProps) => {
  return (
    <div className={c(styles.TabItem, props.tab.active ? styles.active : '')}>
      <button
        className={c(styles.button, props.tab.active ? styles.active : '')}
        onClick={props.onClick}
      >
        {props.tab.label}
      </button>
    </div>
  );
};

type TabsProps = {
  tabs: Tab[];
  onChange: (index: number) => void;
};

const Tabs = (props: TabsProps) => {
  return (
    <nav className={styles.Tabs}>
      {props.tabs.map((tab, index) => (
        <TabItem key={index} tab={tab} onClick={() => props.onChange(index)} />
      ))}
    </nav>
  );
};

export default Tabs;
