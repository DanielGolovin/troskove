import { useQuery } from "@tanstack/react-query";
import { TransactionMonth } from "src/entities/transactions";

export const getYearTransactionsQueryKeyBase = "year-transactions";

export const useGetYearTransactionsQuery = (year: string) => {
  const { data } = useQuery<TransactionMonth[]>({
    queryKey: [getYearTransactionsQueryKeyBase, year],
    queryFn: () => Promise.resolve(mockData[year]),
    // queryFn: () => fetch('/api/categories').then((res) => res.json()),
  });

  return { data };
};

const mockData: Record<string, TransactionMonth[]> = {
  "2024": [
    { name: "March", total: 100 },
    { name: "February", total: 100 },
    { name: "January", total: 100 },
  ],
  "2023": [
    { name: "March", total: 100 },
    { name: "February", total: 100 },
    { name: "January", total: 100 },
  ],
};
