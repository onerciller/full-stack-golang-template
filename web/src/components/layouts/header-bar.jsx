import React from 'react';
import { Layout, Button, Avatar, Dropdown, Space } from 'antd';
import { UserOutlined, BellOutlined } from '@ant-design/icons';
import { HeaderMenuItems } from '@/config';

const { Header } = Layout;

const HeaderBar = ({ collapsed, colorBgContainer }) => {

  return (
    <Header
      style={{
        padding: 0,
        background: colorBgContainer,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        boxShadow: '0 1px 4px rgba(0,21,41,.08)',
        position: 'sticky',
        top: 0,
        zIndex: 999,
      }}
    >
      <Button
        type="text"
        style={{
          fontSize: '16px',
          width: 64,
          height: 64,
        }}
      />
      
      <div style={{ display: 'flex', alignItems: 'center', marginRight: 24 }}>
        <Button icon={<BellOutlined />} style={{ marginRight: 16 }} />
        <Dropdown menu={{ items: HeaderMenuItems }} trigger={['click']} placement="bottomRight">
          <a onClick={(e) => e.preventDefault()}>
            <Space>
              <Avatar icon={<UserOutlined />} />
              {!collapsed && <span style={{ marginLeft: 8 }}>Admin User</span>}
            </Space>
          </a>
        </Dropdown>
      </div>
    </Header>
  );
};

export default HeaderBar; 