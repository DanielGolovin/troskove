import { useQuery } from "@tanstack/react-query";
import { Transaction } from "src/entities/transactions";

export const getMonthTransactionsQueryKeyBase = "month-transactions";

export const useGetMonthTransactionsQuery = (month: string) => {
  const { data } = useQuery<Transaction[]>({
    queryKey: [getMonthTransactionsQueryKeyBase, month],
    queryFn: () => Promise.resolve(mockData[month]),
    // queryFn: () => fetch('/api/categories').then((res) => res.json()),
  });

  return { data };
};

const mockData: Record<string, Transaction[]> = {
  March: [
    { category: "Food", amount: 100, date: "2021-03-01", id: "1" },
    { category: "Food", amount: 100, date: "2021-03-02", id: "2" },
    { category: "Food", amount: 100, date: "2021-03-03", id: "3" },
  ],
  January: [
    { category: "Food", amount: 100, date: "2021-01-01", id: "4" },
    { category: "Food", amount: 100, date: "2021-01-02", id: "5" },
    { category: "Food", amount: 100, date: "2021-01-03", id: "6" },
  ],
  February: [
    { category: "Food", amount: 100, date: "2021-02-01", id: "7" },
    { category: "Food", amount: 100, date: "2021-02-02", id: "8" },
    { category: "Food", amount: 100, date: "2021-02-03", id: "9" },
  ],
};
