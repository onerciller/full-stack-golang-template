import React, { useState } from 'react';
import { Layout, theme } from 'antd';
import SideMenu from './side-menu';
import HeaderBar from './header-bar';
import ContentWrapper from './content-wrapper';

const AdminLayout = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <SideMenu collapsed={collapsed} />
      <Layout>
        <HeaderBar 
          collapsed={collapsed} 
          setCollapsed={setCollapsed} 
          colorBgContainer={colorBgContainer} 
        />
        <ContentWrapper 
          colorBgContainer={colorBgContainer} 
          borderRadiusLG={borderRadiusLG} 
        />
      </Layout>
    </Layout>
  );
};

export default AdminLayout; 