/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions, type InfiniteData  } from "@tanstack/vue-query";
import { useClient } from '../useClient';

export default function usePerpdexLoanV_1() {
  const client = useClient();

  type QueryParamsMethod = typeof client.PerpdexLoanV_1.query.queryParams;
  type QueryParamsData = Awaited<ReturnType<QueryParamsMethod>>["data"];
  const QueryParams = ( options: Partial<UseQueryOptions<QueryParamsData>>) => {
    const key = { type: 'QueryParams',  };    
    return useQuery<QueryParamsData>({ queryKey: [key], queryFn: async () => {
      const res = await client.PerpdexLoanV_1.query.queryParams();
        return res.data;
    }, ...options});
  }
  

  type QueryGetLoanMethod = typeof client.PerpdexLoanV_1.query.queryGetLoan;
  type QueryGetLoanData = Awaited<ReturnType<QueryGetLoanMethod>>["data"];
  const QueryGetLoan = (id: string,  options: Partial<UseQueryOptions<QueryGetLoanData>>) => {
    const key = { type: 'QueryGetLoan',  id };    
    return useQuery<QueryGetLoanData>({ queryKey: [key], queryFn: async () => {
      const { id } = key
      const res = await client.PerpdexLoanV_1.query.queryGetLoan(id);
        return res.data;
    }, ...options});
  }
  
  type QueryListLoanMethod = typeof client.PerpdexLoanV_1.query.queryListLoan;
  type QueryListLoanData = Awaited<ReturnType<QueryListLoanMethod>>["data"] & { pageParam: number };
  const QueryListLoan = (query:  NonNullable<Parameters<QueryListLoanMethod>[0]>, options:  Partial<UseInfiniteQueryOptions<QueryListLoanData, unknown, InfiniteData<QueryListLoanData,number>, Array<string | unknown>, number>> , perPage: number) => {
    const key = { type: 'QueryListLoan', query };    
    return useInfiniteQuery<QueryListLoanData, unknown, InfiniteData<QueryListLoanData,number>, Array<string | unknown>, number>({ queryKey: [key], queryFn: async (context: {pageParam?: number}) => {
      const { pageParam=1 } = context;
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      const res = await client.PerpdexLoanV_1.query.queryListLoan(query ?? undefined);
        return { ...res.data, pageParam }; 
    }, ...options,
      initialPageParam: 1,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryGetLoan,QueryListLoan,
  }
}
