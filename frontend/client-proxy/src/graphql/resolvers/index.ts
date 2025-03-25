import { IResolvers } from "@graphql-tools/utils";

const resolvers: IResolvers = {
  Query: {
    hello: () => "Hello, GraphQL!",
    hello1: () => "<<<Hello, GraphQL!>>>>",
  },
  Mutation: {
    updateMessage: (_: any, { message }: { message: string }) => {
      return `Updated message: ${message}`;
    },
  },
};

export default resolvers;
