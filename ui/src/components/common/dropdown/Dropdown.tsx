import styles from './Dropdown.module.scss';

export type Option = {
  title: string;
  value?: string | number;
};

type Props = {
  options: Option[];
  value?: string;
  label?: string | null;
  name?: string;
  maxWidth?: string | number;
  fullWidth?: boolean;
  testID?: string;
  onChange?: (option: string) => void;
};

function Dropdown({
  fullWidth,
  name,
  label,
  options,
  value,
  maxWidth,
  testID,
  onChange,
}: Props) {
  const renderOptions = () => {
    return options.map(({ title, value }) => (
      <option key={title} value={value}>
        {title}
      </option>
    ));
  };

  return (
    <div className={styles.inputContainer}>
      {label ? (
        <label className={styles.inputLabel} htmlFor={name}>
          {label}
        </label>
      ) : (
        <></>
      )}
      <select
        name={name}
        value={value}
        className={styles.dropdown}
        onChange={(e) => onChange && onChange(e.target.value)}
        style={{ maxWidth, width: fullWidth ? '100%' : 'fit-content' }}
        data-testid={testID}
      >
        {renderOptions()}
      </select>
    </div>
  );
}

export default Dropdown;
