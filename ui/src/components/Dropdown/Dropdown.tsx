import styles from "./Dropdown.module.scss";

export type Option = {
  title: string;
  value?: string | number;
};

type Props = {
  options: Option[];
  value?: string;
  onChange?: (option: string) => void;
};

function Dropdown({ options, value, onChange }: Props) {
  const renderOptions = () => {
    return options.map(({ title, value }) => (
      <option value={value}>{title}</option>
    ));
  };

  return (
    <select
      value={value}
      className={styles.dropdown}
      onChange={(e) => onChange && onChange(e.target.value)}
    >
      {renderOptions()}
    </select>
  );
}

export default Dropdown;
