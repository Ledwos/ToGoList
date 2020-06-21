import React from 'react';
import LogInComp from './LogInComp';
import renderer from 'react-test-renderer';

it('renders LogInComp component', () => {
    const LogInComponent = renderer.create(<LogInComp />).toJSON();
    expect(LogInComponent).toMatchSnapshot();
});