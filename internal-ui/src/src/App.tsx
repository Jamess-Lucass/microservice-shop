import { Suspense, useState } from "react";
import "antd/dist/reset.css";
import {
  Button,
  Center,
  ChakraProvider,
  Spinner,
  useToast,
} from "@chakra-ui/react";
import {
  MutationCache,
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AxiosError } from "axios";
import { ErrorResponse } from "./types/error-response";
import Home from "./pages/home";
import { env } from "./environment";
import Layout from "./components/layout";

function App() {
  const toast = useToast();

  const [queryClient] = useState(
    new QueryClient({
      defaultOptions: {
        queries: {
          retry: false,
          refetchOnMount: false,
          refetchOnWindowFocus: false,
          refetchOnReconnect: false,
        },
      },
      queryCache: new QueryCache({
        onError: (error) => {
          const err = error as AxiosError<ErrorResponse>;

          toast({
            position: "bottom-left",
            title: "Error",
            status: "error",
            description: err.response?.data?.message ?? "An error has occured",
          });
        },
      }),
      mutationCache: new MutationCache({
        onError: (error) => {
          const err = error as AxiosError<ErrorResponse>;

          toast({
            position: "bottom-left",
            title: "Error",
            status: "error",
            description: err.response?.data?.message ?? "An error has occured",
          });
        },
      }),
    })
  );

  return (
    <QueryClientProvider client={queryClient}>
      <ChakraProvider>
        <Layout>
          <BrowserRouter>
            <Suspense
              fallback={
                <Center w="full" h="calc(100vh - 200px)">
                  <Spinner />
                </Center>
              }
            >
              <Routes>
                <Route path="*" element={<Home />} />
              </Routes>
            </Suspense>
          </BrowserRouter>
        </Layout>
      </ChakraProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;
