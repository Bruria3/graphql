import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  split,
  HttpLink,
} from "@apollo/client";
import { getMainDefinition } from "@apollo/client/utilities";
// import { WebSocketLink } from "@apollo/client/link/ws";
import { SSELink } from "@grafbase/apollo-link";

// import { createClient } from 'graphql-sse';

// export const client = createClient({
//   url: 'http://localhost:8080/graphql/subscriptions',
// });

// Web socket
// const wsLink = new WebSocketLink({
//   uri: "ws://localhost:8080/query",
//   options: {
//     reconnect: true,
//   },
// });

const sseLink = new SSELink({
  uri: "http://localhost:8080/api/query",
});

const httpLink = new HttpLink({
  uri: "http://localhost:8080/api/query",
});

const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === "OperationDefinition" &&
      definition.operation === "subscription"
    );
  },
  sseLink,
  httpLink
);

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache(),
});

const ApolloWrapper: React.FC = ({ children }) => (
  <ApolloProvider client={client}>{children}</ApolloProvider>
);

export default ApolloWrapper;
