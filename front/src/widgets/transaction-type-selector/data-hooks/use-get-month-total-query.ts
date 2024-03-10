import { useQuery } from '@tanstack/react-query'
import { TransactionType } from 'src/entities/transactions';

export const getCategoriesQueryKeyBase = 'month-total'

export const useGetMonthTotalQuery = (transactionType: TransactionType) => {
  const { data } = useQuery<number>({
    queryKey: [getCategoriesQueryKeyBase, transactionType],
    queryFn: () => Promise.resolve(mockData[transactionType]),
    // queryFn: () => fetch('/api/categories').then((res) => res.json()),
  });

  return { data };
}

  const mockData = {
    expenses: -1500,
    incomes: 1200,
    mixed: -300,
  };
  