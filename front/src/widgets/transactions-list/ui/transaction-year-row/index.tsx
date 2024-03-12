import { FC, useState } from "react";
import styles from "./index.module.css";
import { TransactionRowItemProps } from "../transaction-row";
import { TransactionMonthRowProps } from "../transaction-month-row";
import { useGetYearTransactionsQuery } from "../../data-hooks/use-get-year-transactions-query";

export type TransactionYearRowProps = {
  amount: number;
  date: string;
  MonthRow: FC<TransactionMonthRowProps>;
  TransactionRow: FC<TransactionRowItemProps>;
};

export const TransactionYearRow: FC<TransactionYearRowProps> = ({
  amount,
  date,
  MonthRow,
  TransactionRow,
}) => {
  const { data = [] } = useGetYearTransactionsQuery(date);

  const [isExpanded, setIsExpanded] = useState(
    () => date === new Date().getFullYear().toString()
  );

  const toggleExpanded = () => {
    setIsExpanded(!isExpanded);
  };

  return (
    <>
      <div onClick={toggleExpanded} className={styles.container}>
        <div>{amount}$</div>
        <div>{date}</div>
        <div>{isExpanded ? ">" : "<"}</div>
      </div>
      {isExpanded && (
        <div className={styles["items-container"]}>
          {data.map(({ name, total }) => (
            <MonthRow
              amount={total}
              date={name}
              key={name}
              TransactionRow={TransactionRow}
            />
          ))}
        </div>
      )}
    </>
  );
};
