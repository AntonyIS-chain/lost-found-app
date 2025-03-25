import { gql } from "apollo-server-express";

const typeDefs = gql`
  type Query {
    hello: String!
    hello1: String!
  }

  type Mutation {
    updateMessage(message: String!): String!
  }
`;

export default typeDefs;
