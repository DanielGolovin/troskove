import { FC, useState } from "react";
import styles from "./index.module.css";
import { TransactionRowItemProps } from "../transaction-row";
import { useGetMonthTransactionsQuery } from "../../data-hooks/use-get-month-transactions-query";

export type TransactionMonthRowProps = {
  amount: number;
  date: string;
  TransactionRow: FC<TransactionRowItemProps>;
};

export const TransactionMonthRow: FC<TransactionMonthRowProps> = ({
  amount,
  date,
  TransactionRow,
}) => {
  const { data = [] } = useGetMonthTransactionsQuery(date);

  const [isExpanded, setIsExpanded] = useState(
    () => date === monthNumberToName(new Date().getMonth())
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
          {data.map(({ category, amount, date, id }) => (
            <TransactionRow
              amount={amount}
              category={category}
              date={date}
              key={id}
              id={id}
            />
          ))}
        </div>
      )}
    </>
  );
};

const monthNumberToName = (month: number) => {
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  return months[month];
};
