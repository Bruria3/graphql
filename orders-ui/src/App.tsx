import "./App.css";
import Orders from "./components/Orders";
import ApolloWrapper from "./services/ApolloClient";

const App: React.FC = () => (
  <ApolloWrapper>
    <Orders />
  </ApolloWrapper>
);

export default App;
