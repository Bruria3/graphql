// import React, { useEffect, useState } from "react";
// import { client } from "../services/SSEClient";
// import { Order, subscriptionOrders } from "../utils/types";
// import { ExecutionResult } from "graphql";
// import { Button, Card } from "@mui/material";

// interface OrdersUpdatedResponse {
//   ordersUpdated: Order[];
// }

// const SSEOrders = () => {
//   const [orders, setOrders] = useState<Order[]>([]);

//   useEffect(() => {
//     const subscribeOrders = async () => {
//       try {
//         const subscription = client.iterate<
//           ExecutionResult<OrdersUpdatedResponse>
//         >({
//           query: subscriptionOrders,
//         });

//         for await (const event of subscription as AsyncIterable<
//           ExecutionResult<OrdersUpdatedResponse>
//         >) {
//           if (event.data && "ordersUpdated" in event.data) {
//             // Update the state with the new orders
//             const newOrders = (event.data as OrdersUpdatedResponse)
//               .ordersUpdated;
//             setOrders((prevOrders) => [...prevOrders, ...newOrders]);
//           }
//         }
//       } catch (error) {
//         console.error("Error subscribing to orders:", error);
//       }
//     };

//     subscribeOrders();

//     // Cleanup function to prevent memory leaks
//     return () => {
//       client.dispose();
//     };
//   }, []);

//   return (
//     <div>
//       <h1>Live Orders</h1>
//       {orders.length === 0 ? (
//         <p>No orders available.</p>
//       ) : (
//         <ul>
//           SSE Orders
//           {orders?.map((order: Order) => (
//             <Card key={order.id}>
//               {order.name} - {order.status} - {order.createdAt}
//               <Button>Update quantity</Button>
//             </Card>
//           ))}
//         </ul>
//       )}
//     </div>
//   );
// };

// export default SSEOrders;
