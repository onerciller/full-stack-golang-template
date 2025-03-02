import React from 'react';
import { Layout } from 'antd';
import { Outlet } from 'react-router-dom';

const { Content } = Layout;

const ContentWrapper = ({ colorBgContainer, borderRadiusLG }) => {
  return (
    <Content
      style={{
        margin: '16px',
        padding: 0,
        minHeight: 280,
        overflow: 'initial',
      }}
    >
      <div
        style={{
          padding: 24,
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <Outlet />
      </div>
    </Content>
  );
};

export default ContentWrapper; 