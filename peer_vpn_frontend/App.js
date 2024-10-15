import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import HomeScreen from './screens/HomeScreen';
import HostScreen from './screens/HostScreen';
import JoinScreen from './screens/JoinScreen';
import ConnectionScreen from './screens/ConnectionScreen';

const Stack = createStackNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="Home">
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Host" component={HostScreen} />
        <Stack.Screen name="Join" component={JoinScreen} />
        <Stack.Screen name="Connection" component={ConnectionScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
