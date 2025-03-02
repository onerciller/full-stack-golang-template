import React, { useState } from 'react';
import { Typography, Table, Tag, Space, Button, message, Spin } from 'antd';
import { useUsers, useDeleteUser } from '@/utils/api/users';

const { Title } = Typography;

const Users = () => {
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
  });
  const [deletingId, setDeletingId] = useState(null);
  
  // Use the useUsers hook to fetch users data
  const { 
    data, 
    isLoading, 
    isError,
    error,
    refetch 
  } = useUsers({ 
    page: pagination.current, 
    limit: pagination.pageSize 
  });
  
  // Use the useDeleteUser hook for deleting users
  const { mutate: deleteUser, isPending: isDeleting } = useDeleteUser();
  
  // Handle deletion
  const handleDelete = (id) => {
    setDeletingId(id);
    deleteUser(id, {
      onSuccess: () => {
        message.success('User deleted successfully');
        setDeletingId(null);
      },
      onError: (err) => {
        message.error(err.message || 'Failed to delete user');
        setDeletingId(null);
      }
    });
  };

  // Handle table pagination change
  const handleTableChange = (pagination) => {
    setPagination({
      current: pagination.current,
      pageSize: pagination.pageSize,
    });
  };
  
  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      render: (text) => <a>{text}</a>,
    },
    {
      title: 'Age',
      dataIndex: 'age',
      key: 'age',
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Role',
      key: 'role',
      dataIndex: 'role',
      render: (role) => (
        <>
          {role === 'admin' ? (
            <Tag color="red">
              {role.toUpperCase()}
            </Tag>
          ) : (
            <Tag color="green">
              {role.toUpperCase()}
            </Tag>
          )}
        </>
      ),
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button type="link">Edit</Button>
          <Button 
            type="link" 
            danger 
            loading={isDeleting && record.key === deletingId}
            onClick={() => handleDelete(record.key)}
          >
            Delete
          </Button>
        </Space>
      ),
    },
  ];

  // Fallback data for development if the API is not yet available
  const fallbackData = [
    {
      key: '1',
      name: 'John Brown',
      age: 32,
      email: 'john@example.com',
      role: 'admin',
    },
    {
      key: '2',
      name: 'Jim Green',
      age: 42,
      email: 'jim@example.com',
      role: 'user',
    },
    {
      key: '3',
      name: 'Joe Black',
      age: 32,
      email: 'joe@example.com',
      role: 'user',
    },
  ];

  // Display loading or error states
  if (isLoading) {
    return (
      <div style={{ textAlign: 'center', padding: '50px' }}>
        <Spin size="large" />
        <p>Loading users...</p>
      </div>
    );
  }

  if (isError) {
    return (
      <div style={{ textAlign: 'center', padding: '50px' }}>
        <p>Error loading users: {error?.message || 'Unknown error'}</p>
        <Button onClick={() => refetch()}>Try Again</Button>
      </div>
    );
  }

  // Use the API data or fallback if in development without API
  const usersData = data?.users || fallbackData;
  const total = data?.total || usersData.length;

  return (
    <div>
      <Title level={2}>Users</Title>
      <Table 
        columns={columns} 
        dataSource={usersData}
        pagination={{
          ...pagination,
          total,
          showSizeChanger: true,
        }}
        onChange={handleTableChange}
        loading={isLoading}
      />
    </div>
  );
};

export default Users; 