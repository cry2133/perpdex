/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions, type InfiniteData  } from "@tanstack/vue-query";
import { useClient } from '../useClient';

export default function usePerpdexPerpdexV_1() {
  const client = useClient();

  type QueryParamsMethod = typeof client.PerpdexPerpdexV_1.query.queryParams;
  type QueryParamsData = Awaited<ReturnType<QueryParamsMethod>>["data"];
  const QueryParams = ( options: Partial<UseQueryOptions<QueryParamsData>>) => {
    const key = { type: 'QueryParams',  };    
    return useQuery<QueryParamsData>({ queryKey: [key], queryFn: async () => {
      const res = await client.PerpdexPerpdexV_1.query.queryParams();
        return res.data;
    }, ...options});
  }
  

  type QueryShowPriceMethod = typeof client.PerpdexPerpdexV_1.query.queryShowPrice;
  type QueryShowPriceData = Awaited<ReturnType<QueryShowPriceMethod>>["data"];
  const QueryShowPrice = (id: string,  options: Partial<UseQueryOptions<QueryShowPriceData>>) => {
    const key = { type: 'QueryShowPrice',  id };    
    return useQuery<QueryShowPriceData>({ queryKey: [key], queryFn: async () => {
      const { id } = key
      const res = await client.PerpdexPerpdexV_1.query.queryShowPrice(id);
        return res.data;
    }, ...options});
  }
  

  type QueryShowSymbolPriceMethod = typeof client.PerpdexPerpdexV_1.query.queryShowSymbolPrice;
  type QueryShowSymbolPriceData = Awaited<ReturnType<QueryShowSymbolPriceMethod>>["data"];
  const QueryShowSymbolPrice = (symbol: string,  options: Partial<UseQueryOptions<QueryShowSymbolPriceData>>) => {
    const key = { type: 'QueryShowSymbolPrice',  symbol };    
    return useQuery<QueryShowSymbolPriceData>({ queryKey: [key], queryFn: async () => {
      const { symbol } = key
      const res = await client.PerpdexPerpdexV_1.query.queryShowSymbolPrice(symbol);
        return res.data;
    }, ...options});
  }
  
  type QueryListPriceMethod = typeof client.PerpdexPerpdexV_1.query.queryListPrice;
  type QueryListPriceData = Awaited<ReturnType<QueryListPriceMethod>>["data"] & { pageParam: number };
  const QueryListPrice = (query:  NonNullable<Parameters<QueryListPriceMethod>[0]>, options:  Partial<UseInfiniteQueryOptions<QueryListPriceData, unknown, InfiniteData<QueryListPriceData,number>, Array<string | unknown>, number>> , perPage: number) => {
    const key = { type: 'QueryListPrice', query };    
    return useInfiniteQuery<QueryListPriceData, unknown, InfiniteData<QueryListPriceData,number>, Array<string | unknown>, number>({ queryKey: [key], queryFn: async (context: {pageParam?: number}) => {
      const { pageParam=1 } = context;
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      const res = await client.PerpdexPerpdexV_1.query.queryListPrice(query ?? undefined);
        return { ...res.data, pageParam }; 
    }, ...options,
      initialPageParam: 1,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryShowPrice,QueryShowSymbolPrice,QueryListPrice,
  }
}
